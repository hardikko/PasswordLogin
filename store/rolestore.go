package store

import (
	"context"
	"learngo/models"
	"learngo/settings"
	"learngo/utils/faulterr"

	"github.com/jackc/pgx/v4"
)

func RoleListStore(ctx context.Context) ([]models.Role, *faulterr.FaultErr) {
	result := []models.Role{}
	obj := models.Role{}
	errMsg := "error when trying to get roles"

	queryStmt := `SELECT * FROM roles`
	rows, err := settings.DBClient.Query(ctx, queryStmt)
	if err != nil {
		return nil, faulterr.NewPostgresError(err, errMsg)
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(
			&obj.ID,
			&obj.Code,
			&obj.OrgID,
			&obj.DepartID,
			&obj.Name,
			&obj.Permissions,
			&obj.IsManagement,
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
func RoleGetByIDStore(ctx context.Context, id int64) (*models.Role, *faulterr.FaultErr) {
	obj := models.Role{}
	errMsg := "error when trying to get role by id"

	queryStmt := `
		SELECT * FROM roles
		WHERE roles.id = $1
	`
	row := settings.DBClient.QueryRow(ctx, queryStmt, id)
	if err := row.Scan(
		&obj.ID,
		&obj.Code,
		&obj.OrgID,
		&obj.DepartID,
		&obj.Name,
		&obj.Permissions,
		&obj.IsManagement,
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

func RoleInsertStore(ctx context.Context, tx pgx.Tx, arg models.Role) (*models.Role, *faulterr.FaultErr) {
	obj := &models.Role{}
	errMsg := "error when trying to insert role"

	queryStmt := `
	INSERT INTO
	roles(
		code,
		org_id,
		depart_id,
		name,
		permissions,
		is_management,
		status,
		created_at,
		updated_at

	)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	RETURNING *
	`

	row := tx.QueryRow(ctx, queryStmt,
		&arg.Code,
		&arg.OrgID,
		&arg.DepartID,
		&arg.Name,
		&arg.Permissions,
		&arg.IsManagement,
		&arg.Status,
		&arg.CreatedAt,
		&arg.UpdatedAt,
	)

	if err := row.Scan(
		&obj.ID,
		&obj.Code,
		&obj.OrgID,
		&obj.DepartID,
		&obj.Name,
		&obj.Permissions,
		&obj.IsManagement,
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

func UpdateRoleStore(ctx context.Context, tx pgx.Tx, arg models.Role) *faulterr.FaultErr {
	errMsg := "error when trying to update role"

	queryStmt := `
	UPDATE roles
	SET
		name=$1,
		status=$2
	WHERE id=$3
	`

	_, err := tx.Exec(ctx, queryStmt,
		&arg.Name,
		&arg.Status,
		&arg.ID,
	)
	if err != nil {
		return faulterr.NewPostgresError(err, errMsg)
	}
	return nil
}
