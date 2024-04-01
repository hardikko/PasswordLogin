package store

import (
	"context"
	"learngo/models"
	"learngo/settings"
	"learngo/utils/faulterr"

	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v4"
)

func AuthGetByIDStore(ctx context.Context, id int64) (*models.Authtoken, *faulterr.FaultErr) {
	obj := models.Authtoken{}
	errMsg := "error when trying to get AuthToken by id"

	queryStmt := `
		SELECT * FROM auth_sessions
		WHERE auth_sessions.id = $1
	`
	row := settings.DBClient.QueryRow(ctx, queryStmt, id)
	if err := row.Scan(
		&obj.ID,
		&obj.UserID,
		&obj.Token,
		&obj.IsValid,
		&obj.ExpireAt,
		&obj.CreatedAt,
		&obj.UpdatedAt,
	); err != nil {
		return nil, faulterr.NewPostgresError(err, errMsg)
	}

	return &obj, nil
}

func AuthInsertStore(ctx context.Context, tx pgx.Tx, arg models.Authtoken) (*models.Authtoken, *faulterr.FaultErr) {
	obj := &models.Authtoken{}
	errMsg := "error when trying to insert AuthToken"

	queryStmt := `
	INSERT INTO
	auth_sessions(
		user_id,
		token,
		is_valid,
		expire_at,
		created_at,
		updated_at

	)
	VALUES ($1, $2, $3, $4, $5, $6)
	RETURNING *
	`

	row := tx.QueryRow(ctx, queryStmt,
		&arg.UserID,
		&arg.Token,
		&arg.IsValid,
		&arg.ExpireAt,
		&arg.CreatedAt,
		&arg.UpdatedAt,
	)

	if err := row.Scan(
		&obj.ID,
		&obj.UserID,
		&obj.Token,
		&obj.IsValid,
		&obj.ExpireAt,
		&obj.CreatedAt,
		&obj.UpdatedAt,
	); err != nil {
		return nil, faulterr.NewPostgresError(err, errMsg)
	}

	return obj, nil
}

func AuthGetByTokenStore(ctx context.Context, token uuid.UUID) (*models.Authtoken, *faulterr.FaultErr) {
	obj := models.Authtoken{}
	errMsg := "error when trying to get auth session by token"

	queryStmt := `
	SELECT * FROM auth_sessions
	WHERE auth_sessions.token = $1
	`
	row := settings.DBClient.QueryRow(ctx, queryStmt, token)
	if err := row.Scan(
		&obj.ID,
		&obj.UserID,
		&obj.Token,
		&obj.IsValid,
		&obj.ExpireAt,
		&obj.CreatedAt,
		&obj.UpdatedAt,
	); err != nil {
		return nil, faulterr.NewPostgresError(err, errMsg)
	}

	return &obj, nil
}
