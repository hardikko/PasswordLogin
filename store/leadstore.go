package store

import (
	"context"

	"learngo/models"
	"learngo/settings"
	"learngo/utils/faulterr"

	"github.com/jackc/pgx/v4"
)

func LeadListStore(ctx context.Context) ([]models.Lead, *faulterr.FaultErr) {
	result := []models.Lead{}
	obj := models.Lead{}
	errMsg := "error when trying to get lead"

	queryStmt := `SELECT * FROM leads`
	rows, err := settings.DBClient.Query(ctx, queryStmt)
	if err != nil {
		return nil, faulterr.NewPostgresError(err, errMsg)
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(
			&obj.ID,
			&obj.UserID,
			&obj.FirstName,
			&obj.LastName,
			&obj.Address,
			&obj.Phone,
			&obj.Email,
			&obj.Occupation,
			&obj.Company,
			&obj.Status,
			&obj.CreatedAt,
			&obj.UpdatedAt,
		); err != nil {
			return nil, faulterr.NewPostgresError(err, errMsg)
		}
		result = append(result, obj)
	}

	return result, nil
}
func LeadGetByIDStore(ctx context.Context, id int64) (*models.Lead, *faulterr.FaultErr) {
	obj := models.Lead{}
	errMsg := "error when trying to get lead by id"

	queryStmt := `
		SELECT * FROM leads
		WHERE leads.id = $1
	`
	row := settings.DBClient.QueryRow(ctx, queryStmt, id)
	if err := row.Scan(
		&obj.ID,
		&obj.UserID,
		&obj.FirstName,
		&obj.LastName,
		&obj.Address,
		&obj.Phone,
		&obj.Email,
		&obj.Occupation,
		&obj.Company,
		&obj.Status,
		&obj.CreatedAt,
		&obj.UpdatedAt,
	); err != nil {
		return nil, faulterr.NewPostgresError(err, errMsg)
	}

	return &obj, nil
}

func LeadInsertStore(ctx context.Context, tx pgx.Tx, arg models.Lead) (*models.Lead, *faulterr.FaultErr) {
	obj := &models.Lead{}
	errMsg := "error when trying to insert leads"

	queryStmt := `
	INSERT INTO
	leads(
		user_id,
		firstname,
		lastname,
		address,
		phone,
		email,
		occupation,
		company,
		status,
		created_at,
		updated_at

	)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
	RETURNING *
	`

	row := tx.QueryRow(ctx, queryStmt,
		&arg.UserID,
		&arg.FirstName,
		&arg.LastName,
		&arg.Address,
		&arg.Phone,
		&arg.Email,
		&arg.Occupation,
		&arg.Company,
		&arg.Status,
		&arg.CreatedAt,
		&arg.UpdatedAt,
	)

	if err := row.Scan(
		&obj.UserID,
		&obj.ID,
		&obj.FirstName,
		&obj.LastName,
		&obj.Address,
		&obj.Phone,
		&obj.Email,
		&obj.Occupation,
		&obj.Company,
		&obj.Status,
		&obj.CreatedAt,
		&obj.UpdatedAt,
	); err != nil {
		return nil, faulterr.NewPostgresError(err, errMsg)
	}

	return obj, nil
}

func UpdateLeadStore(ctx context.Context, tx pgx.Tx, arg models.Lead) *faulterr.FaultErr {
	errMsg := "error when trying to update leads"

	queryStmt := `
	UPDATE leads
	SET
		firstname=$1,
		lastname=$2,
		address=$3,
		phone=$4,
		email=$5,
		occupation=$6,
		status=$7
	WHERE id=$8
	`

	_, err := tx.Exec(ctx, queryStmt,
		&arg.FirstName,
		&arg.LastName,
		&arg.Address,
		&arg.Phone,
		&arg.Email,
		&arg.Occupation,
		&arg.Status,
		&arg.ID,
	)
	if err != nil {
		return faulterr.NewPostgresError(err, errMsg)
	}
	return nil
}
