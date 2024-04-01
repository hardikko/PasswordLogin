package store

import (
	"context"
	"fmt"
	"learngo/models"
	"learngo/settings"
	"learngo/utils/faulterr"

	"github.com/jackc/pgx/v4"
)

func UserListStore(ctx context.Context) ([]models.User, *faulterr.FaultErr) {
	result := []models.User{}
	obj := models.User{}
	errMsg := "error when trying to get users"

	queryStmt := `SELECT * FROM users`
	rows, err := settings.DBClient.Query(ctx, queryStmt)
	if err != nil {
		return nil, faulterr.NewPostgresError(err, errMsg)
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(
			&obj.ID,
			&obj.Code,
			&obj.FirstName,
			&obj.LastName,
			&obj.Email,
			&obj.Phone,
			&obj.PasswordHash,
			&obj.IsAdmin,
			&obj.OrgID,
			&obj.RoleID,
			&obj.Status,
			&obj.IsFinal,
			&obj.IsArchived,
			&obj.CreatedAt,
			&obj.UpdatedAt,
		); err != nil {
			return nil, faulterr.NewPostgresError(err, errMsg)
		}
		result = append(result, obj)
	}

	return result, nil
}
func UserGetByIDStore(ctx context.Context, id int64) (*models.User, *faulterr.FaultErr) {
	obj := models.User{}
	errMsg := "error when trying to get user by id"

	queryStmt := `
		SELECT * FROM users
		WHERE users.id = $1
	`
	row := settings.DBClient.QueryRow(ctx, queryStmt, id)
	if err := row.Scan(
		&obj.ID,
		&obj.Code,
		&obj.FirstName,
		&obj.LastName,
		&obj.Email,
		&obj.Phone,
		&obj.PasswordHash,
		&obj.IsAdmin,
		&obj.OrgID,
		&obj.RoleID,
		&obj.Status,
		&obj.IsFinal,
		&obj.IsArchived,
		&obj.CreatedAt,
		&obj.UpdatedAt,
	); err != nil {
		return nil, faulterr.NewPostgresError(err, errMsg)
	}

	return &obj, nil
}

func UserInsertStore(ctx context.Context, tx pgx.Tx, arg models.User) (*models.User, *faulterr.FaultErr) {
	obj := &models.User{}
	errMsg := "error when trying to insert user"

	queryStmt := `
	INSERT INTO
	users(
		code,
		first_name,
		last_name,
		email,
		phone,
		password_hash,
		is_admin,
		org_id,
		role_id,
		status,
		is_final,
		is_archived

	)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
	RETURNING *
	`

	row := tx.QueryRow(ctx, queryStmt,
		&arg.Code,
		&arg.FirstName,
		&arg.LastName,
		&arg.Email,
		&arg.Phone,
		&arg.PasswordHash,
		&arg.IsAdmin,
		&arg.OrgID,
		&arg.RoleID,
		&arg.Status,
		&arg.IsFinal,
		&arg.IsArchived,
	)

	fmt.Println(arg)
	fmt.Println(arg.Status)
	fmt.Println(obj)
	fmt.Println(obj.Status)

	if err := row.Scan(
		&obj.ID,
		&obj.Code,
		&obj.FirstName,
		&obj.LastName,
		&obj.Email,
		&obj.Phone,
		&obj.PasswordHash,
		&obj.IsAdmin,
		&obj.OrgID,
		&obj.RoleID,
		&obj.Status,
		&obj.IsFinal,
		&obj.IsArchived,
		&obj.CreatedAt,
		&obj.UpdatedAt,
	); err != nil {
		return nil, faulterr.NewPostgresError(err, errMsg)
	}

	return obj, nil
}

func GetUsersByEmailStore(ctx context.Context, email string) (*models.User, *faulterr.FaultErr) {
	obj := models.User{}
	errMsg := fmt.Sprintf("error when trying to get user by email - %s", email)

	queryStmt := `
		SELECT * FROM users
		WHERE users.email = $1
	`
	row := settings.DBClient.QueryRow(ctx, queryStmt, email)
	if err := row.Scan(
		&obj.ID,
		&obj.Code,
		&obj.FirstName,
		&obj.LastName,
		&obj.Email,
		&obj.Phone,
		&obj.PasswordHash,
		&obj.IsAdmin,
		&obj.OrgID,
		&obj.RoleID,
		&obj.Status,
		&obj.IsFinal,
		&obj.IsArchived,
		&obj.CreatedAt,
		&obj.UpdatedAt,
	); err != nil {
		return nil, faulterr.NewPostgresError(err, errMsg)
	}

	return &obj, nil
}

func GetUsersByPhoneStore(ctx context.Context, phone string) (*models.User, *faulterr.FaultErr) {
	obj := models.User{}
	errMsg := fmt.Sprintf("error when trying to get user by phone - %s", phone)

	queryStmt := `
		SELECT * FROM users
		WHERE users.phone = $1
	`
	row := settings.DBClient.QueryRow(ctx, queryStmt, phone)
	if err := row.Scan(
		&obj.ID,
		&obj.Code,
		&obj.FirstName,
		&obj.LastName,
		&obj.Email,
		&obj.Phone,
		&obj.PasswordHash,
		&obj.IsAdmin,
		&obj.OrgID,
		&obj.RoleID,
		&obj.Status,
		&obj.IsFinal,
		&obj.IsArchived,
		&obj.CreatedAt,
		&obj.UpdatedAt,
	); err != nil {
		return nil, faulterr.NewPostgresError(err, errMsg)
	}

	return &obj, nil
}
func UpdateUserStore(ctx context.Context, tx pgx.Tx, arg models.User) *faulterr.FaultErr {
	errMsg := "error when trying to update user"

	queryStmt := `
	UPDATE users
	SET
		first_name=$1,
		last_name=$2,
		email=$3,
		phone=$4
	WHERE id=$5
	`

	_, err := tx.Exec(ctx, queryStmt,
		&arg.FirstName,
		&arg.LastName,
		&arg.Email,
		&arg.Phone,
		&arg.ID,
	)
	if err != nil {
		return faulterr.NewPostgresError(err, errMsg)
	}
	return nil
}
