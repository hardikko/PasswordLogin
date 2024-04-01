package services

import (
	"context"
	"learngo/helpers"
	"learngo/models"
	"learngo/settings"
	"learngo/store"
	"learngo/utils/faulterr"
	"net/http"
	"time"

	"github.com/jackc/pgx"
)

func OtpService(ctx context.Context, id int64) (*models.Otp, *faulterr.FaultErr) {
	obj, err := store.OtpGetByIDStore(ctx, id)
	if err != nil {
		return nil, err
	}

	return obj, nil
}

func OtpCreateService(ctx context.Context, req models.Otp) (*models.Otp, *faulterr.FaultErr) {
	// connect to dbstore
	arg := &models.Otp{}
	tx, err := settings.BeginTx(ctx)
	if err != nil {
		return nil, err
	}
	defer settings.RollbackTx(ctx, tx)

	obj, err := store.OtpInsertStore(ctx, tx, arg)
	if err != nil {
		return nil, err
	}

	if err := settings.CommitTx(ctx, tx); err != nil {
		return nil, err
	}

	return obj, nil
}

func GetOTPService(ctx context.Context, tx pgx.Tx, req *models.OTPRequest) (*string, *faulterr.FaultErr) {
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
			return nil, err
		}
		user = obj
	}

	//generate OTP
	otp, err := Create(ctx, tx, user.ID)
	if err != nil {
		return nil, err
	}
	return &otp.Token, nil
}

func Create(ctx context.Context, tx pgx.Tx, userID int64) (*models.Otp, *faulterr.FaultErr) {
	return store.OtpInsertStore(ctx, tx, construct(userID))
}

func construct(userID int64) *models.Otp {
	return &models.Otp{
		UserID:    userID,
		Token:     helpers.GenerateRandomString(5),
		IsValid:   true,
		ExpiresAt: time.Now().Add(time.Minute * time.Duration(5)), // token will expire after 5 minutes
	}
}
