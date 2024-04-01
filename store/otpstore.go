package store

import (
	"context"
	"fmt"
	"learngo/models"
	"learngo/settings"
	"learngo/utils/faulterr"

	"github.com/jackc/pgx/v4"
)

func OtpGetByIDStore(ctx context.Context, id int64) (*models.Otp, *faulterr.FaultErr) {
	obj := models.Otp{}
	errMsg := "error when trying to get Otp by id"

	queryStmt := `
		SELECT * FROM otp_sessions
		WHERE otp_sessions.id = $1
	`
	row := settings.DBClient.QueryRow(ctx, queryStmt, id)
	if err := row.Scan(
		&obj.ID,
		&obj.UserID,
		&obj.Token,
		&obj.IsValid,
		&obj.ExpiresAt,
		&obj.CreatedAt,
		&obj.UpdatedAt,
	); err != nil {
		return nil, faulterr.NewPostgresError(err, errMsg)
	}

	return &obj, nil
}

func OtpInsertStore(ctx context.Context, tx pgx.Tx, arg *models.Otp) (*models.Otp, *faulterr.FaultErr) {
	obj := &models.Otp{}
	errMsg := "error when trying to insert otp session"

	fmt.Println("OTP store")

	queryStmt := `
	INSERT INTO
	otp_sessions(
		user_id,
		token,
		is_valid,
		expires_at

	)
	VALUES ($1, $2, $3, $4)
	RETURNING *
	`

	row := tx.QueryRow(ctx, queryStmt,
		&arg.UserID,
		&arg.Token,
		&arg.IsValid,
		&arg.ExpiresAt,
	)

	if err := row.Scan(
		&obj.ID,
		&obj.UserID,
		&obj.Token,
		&obj.IsValid,
		&obj.ExpiresAt,
		&obj.CreatedAt,
		&obj.UpdatedAt,
	); err != nil {
		return nil, faulterr.NewPostgresError(err, errMsg)
	}

	fmt.Println("OTP store insert success")

	return obj, nil
}

// OtpGetByTokenStore gets otp session by code from db!
func OtpGetByTokenStore(ctx context.Context, token string) (*models.Otp, *faulterr.FaultErr) {
	obj := models.Otp{}
	errMsg := "error when trying to get otp session by token"

	queryStmt := `
	SELECT * FROM otp_sessions
	WHERE otp_sessions.token = $1
	`

	row := settings.DBClient.QueryRow(ctx, queryStmt, token)
	if err := row.Scan(
		&obj.ID,
		&obj.UserID,
		&obj.Token,
		&obj.IsValid,
		&obj.ExpiresAt,
		&obj.CreatedAt,
		&obj.UpdatedAt,
	); err != nil {
		return nil, faulterr.NewPostgresError(err, errMsg)
	}

	return &obj, nil
}

// Update User
func UpdateOtpStore(ctx context.Context, tx pgx.Tx, arg *models.Otp) *faulterr.FaultErr {
	errMsg := "error when trying to update user"

	queryStmt := `
	UPDATE otp_sessions
	SET
		is_valid=$1,
	WHERE id=$2
	`

	_, err := tx.Exec(ctx, queryStmt,
		&arg.IsValid,
		&arg.ID,
	)
	if err != nil {
		return faulterr.NewPostgresError(err, errMsg)
	}
	return nil
}
