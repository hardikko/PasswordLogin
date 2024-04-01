package store

import (
	"context"
	"learngo/models"
	"learngo/settings"
	"learngo/utils/faulterr"

	"github.com/jackc/pgx/v4"
)

func DepartListStore(ctx context.Context) ([]models.Department, *faulterr.FaultErr) {
	result := []models.Department{}
	obj := models.Department{}
	errMsg := "error when trying to get departments"

	queryStmt := `SELECT * FROM departments`
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
			&obj.Name,
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
func DepartGetByIDStore(ctx context.Context, id int64) (*models.Department, *faulterr.FaultErr) {
	obj := models.Department{}
	errMsg := "error when trying to get departments by id"

	queryStmt := `
		SELECT * FROM departments
		WHERE departments.id = $1
	`
	row := settings.DBClient.QueryRow(ctx, queryStmt, id)
	if err := row.Scan(
		&obj.ID,
		&obj.Code,
		&obj.OrgID,
		&obj.Name,
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

func DepartInsertStore(ctx context.Context, tx pgx.Tx, arg models.Department) (*models.Department, *faulterr.FaultErr) {
	obj := &models.Department{}
	errMsg := "error when trying to insert departments "

	queryStmt := `
	INSERT INTO
	departments(
		code,
		org_id,
		name,
		status,
		is_final,
		is_archived,
		created_at,
		updated_at

	)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	RETURNING *
	`

	row := tx.QueryRow(ctx, queryStmt,
		&arg.Code,
		&arg.OrgID,
		&arg.Name,
		&arg.Status,
		&arg.IsFinal,
		&arg.IsArchived,
		&arg.CreatedAt,
		&arg.UpdatedAt,
	)

	if err := row.Scan(
		&obj.ID,
		&obj.Code,
		&obj.OrgID,
		&obj.Name,
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

func UpdateDepartStore(ctx context.Context, tx pgx.Tx, arg models.Department) *faulterr.FaultErr {
	errMsg := "error when trying to update departments"

	queryStmt := `
	UPDATE departments
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
