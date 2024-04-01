package store

import (
	"context"
	"learngo/models"
	"learngo/settings"
	"learngo/utils/faulterr"

	"github.com/jackc/pgx/v4"
)

func OrgListStore(ctx context.Context) ([]models.Organization, *faulterr.FaultErr) {
	result := []models.Organization{}
	obj := models.Organization{}
	errMsg := "error when trying to get organization"

	queryStmt := `SELECT * FROM organizations`
	rows, err := settings.DBClient.Query(ctx, queryStmt)
	if err != nil {
		return nil, faulterr.NewPostgresError(err, errMsg)
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(
			&obj.ID,
			&obj.Code,
			&obj.Name,
			&obj.Website,
			&obj.Sector,
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
func OrgGetByIDStore(ctx context.Context, id int64) (*models.Organization, *faulterr.FaultErr) {
	obj := models.Organization{}
	errMsg := "error when trying to get organization by id"

	queryStmt := `
		SELECT * FROM organizations
		WHERE organizations.id = $1
	`
	row := settings.DBClient.QueryRow(ctx, queryStmt, id)
	if err := row.Scan(
		&obj.ID,
		&obj.Code,
		&obj.Name,
		&obj.Website,
		&obj.Sector,
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

func OrgInsertStore(ctx context.Context, tx pgx.Tx, arg models.Organization) (*models.Organization, *faulterr.FaultErr) {
	obj := &models.Organization{}
	errMsg := "error when trying to insert organization"

	queryStmt := `
	INSERT INTO
	organizations(
		code,
		name,
		website,
		sector,
		status,
		created_at,
		updated_at

	)
	VALUES ($1, $2, $3, $4, $5, $6, $7)
	RETURNING *
	`

	row := tx.QueryRow(ctx, queryStmt,
		&arg.Code,
		&arg.Name,
		&arg.Website,
		&arg.Sector,
		&arg.Status,
		&arg.CreatedAt,
		&arg.UpdatedAt,
	)

	if err := row.Scan(
		&obj.ID,
		&obj.Code,
		&obj.Name,
		&obj.Website,
		&obj.Sector,
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

func UpdateOrgStore(ctx context.Context, tx pgx.Tx, arg models.Organization) *faulterr.FaultErr {
	errMsg := "error when trying to update organization"

	queryStmt := `
	UPDATE organizations
	SET
		name=$1,
		website=$2,
		sector=$3,
		status=$4
	WHERE id=$5
	`

	_, err := tx.Exec(ctx, queryStmt,
		&arg.Name,
		&arg.Website,
		&arg.Sector,
		&arg.Status,
		&arg.ID,
	)
	if err != nil {
		return faulterr.NewPostgresError(err, errMsg)
	}
	return nil
}
