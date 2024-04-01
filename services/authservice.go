package services

import (
	"context"
	"fmt"
	"learngo/helpers"
	"learngo/models"
	"learngo/settings"
	"learngo/store"
	"learngo/utils/faulterr"
	"net/http"
	"time"

	"github.com/gofrs/uuid"
	"github.com/jackc/pgx"
)

func LoginCreateService(ctx context.Context, req models.Login) (*models.Author, *faulterr.FaultErr) {
	// connect to dbstore
	tx, err := settings.BeginTx(ctx)
	if err != nil {
		return nil, err
	}
	defer settings.RollbackTx(ctx, tx)

	//Check if user exists in db
	obj, err := store.GetUsersByEmailStore(ctx, req.Email.String)
	if err != nil {
		return nil, err
	}

	req.PasswordHash = helpers.GetMd5(req.PasswordHash)

	// Compare User pass with req pass
	if req.PasswordHash != obj.PasswordHash {
		return nil, err
	}
	tokenUUID, err := uuid.DefaultGenerator.NewV4()
	if err != nil {
		return nil, err
	}

	res := models.Authtoken{}
	res.UserID = obj.ID
	res.Token = tokenUUID
	res.ExpireAt = time.Now().Add(24 * time.Hour)

	auth, err := store.AuthInsertStore(ctx, tx, res)
	if err != nil {
		return nil, err
	}

	if req.Email.String != obj.Email {
		return nil, err
	}
	if err := settings.CommitTx(ctx, tx); err != nil {
		return nil, err
	}
	author := &models.Author{}
	author.UserID = auth.UserID
	author.Name = obj.FirstName + obj.LastName
	author.Token = auth.Token

	return author, nil
}

func Login(ctx context.Context, tx pgx.Tx, req *models.Login) (*models.Auther, *faulterr.FaultErr) {
	user := &models.User{}

	if req.Email.Valid && req.Email.String != "" {
		obj, err := store.GetUsersByEmailStore(ctx, req.Email.String)
		if err != nil {
			if err.Status == http.StatusNotFound {
				return nil, faulterr.NewFrobiddenError("no user found with given email")
			}
			return nil, err
		}
		user = obj

	}
	if req.Phone.Valid && req.Email.String != "" {
		obj, err := store.GetUsersByPhoneStore(ctx, req.Phone.String)
		if err != nil {
			if err.Status == http.StatusNotFound {
				return nil, faulterr.NewFrobiddenError("no user found with given phone")
			}
			return nil, err
		}
		user = obj
	}

	// get otp from db and validate
	otp, err := store.OtpGetByTokenStore(ctx, req.OTP)
	errMsg := "error when trying to get otp from db"
	if err != nil {
		if err.Status == http.StatusNotFound {
			return nil, faulterr.NewFrobiddenError("otp is invalid")
		}
		return nil, err
	}
	if !otp.IsValid {
		return nil, faulterr.NewUnauthorizedError("	otp is not valid")
	}
	if err := helpers.ValidateTokenExpiry(otp.ExpiresAt); err != nil {
		return nil, faulterr.NewPostgresError(err, errMsg)
	}

	// update otp validity
	otp.IsValid = false
	if err := store.UpdateOtpStore(ctx, tx, *otp); err != nil {
		return nil, err
	}

	// generate auth session token
	authSession, err := AuthCreate(ctx, tx, user.ID)
	if err != nil {
		return nil, err
	}
	return getAuther(user, authSession.Token), nil
}

func GetAutherByTokenService(ctx context.Context, token uuid.UUID) (*models.Auther, *faulterr.FaultErr) {
	errMsg := "error when trying to get auther by token"
	authSession, err := store.AuthGetByTokenStore(ctx, token)
	if err != nil {
		return nil, err
	}
	if !authSession.IsValid {
		return nil, faulterr.NewUnauthorizedError("token is not valid")
	}
	if err := helpers.ValidateTokenExpiry(authSession.ExpireAt); err != nil {
		return nil, faulterr.NewPostgresError(err, errMsg)
	}

	// get user
	user, err := store.UserGetByIDStore(ctx, authSession.UserID)
	if err != nil {
		return nil, err
	}
	return getAuther(user, token), nil
}

// helpers
func getAuther(u *models.User, token uuid.UUID) *models.Auther {
	name := fmt.Sprintf("%s %s", u.FirstName, u.LastName)
	return &models.Auther{
		ID:           u.ID,
		Name:         name,
		IsAdmin:      u.IsAdmin,
		OrgID:        u.OrgID,
		RoleID:       u.RoleID,
		SessionToken: token,
	}
}

func AuthCreate(ctx context.Context, tx pgx.Tx, userID int64) (*models.Authtoken, *faulterr.FaultErr) {
	obj, err := Authconstruct(userID)
	if err != nil {
		return nil, err
	}
	return store.AuthInsertStore(ctx, tx, *obj)
}

func Authconstruct(userID int64) (*models.Authtoken, *faulterr.FaultErr) {
	token, err := helpers.GenerateUID()
	if err != nil {
		return nil, err
	}

	obj := &models.Authtoken{
		UserID:   userID,
		Token:    *token,
		IsValid:  true,
		ExpireAt: time.Now().Add(time.Hour * time.Duration(720)), // token will expire after 5 minutes
	}

	return obj, nil
}
