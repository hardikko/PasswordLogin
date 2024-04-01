package services

import (
	"context"
	"fmt"
	"learngo/models"
	"learngo/settings"
	"learngo/store"
	"learngo/utils/faulterr"
)

func OrgService(ctx context.Context, id int64) (*models.Organization, *faulterr.FaultErr) {
	obj, err := store.OrgGetByIDStore(ctx, id)
	if err != nil {
		return nil, err
	}

	return obj, nil
}

func OrgCreateService(ctx context.Context, req models.Organization) (*models.Organization, *faulterr.FaultErr) {
	//Get all Organization list
	orgs, err := store.OrgListStore(ctx)

	if err != nil {
		return nil, err
	}
	code := fmt.Sprintf("ORG%03d", len(orgs)+1)
	req.Code = code

	// connect to dbstore
	tx, err := settings.BeginTx(ctx)
	if err != nil {
		return nil, err
	}
	defer settings.RollbackTx(ctx, tx)

	obj, err := store.OrgInsertStore(ctx, tx, req)
	if err != nil {
		return nil, err
	}

	if err := settings.CommitTx(ctx, tx); err != nil {
		return nil, err
	}

	return obj, nil
}

func OrgListService(ctx context.Context) ([]models.Organization, *faulterr.FaultErr) {
	result, err := store.OrgListStore(ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func OrgUpdateService(ctx context.Context, id int64, req models.Organization) (*models.Organization, *faulterr.FaultErr) {
	// get organization by id from db
	org, err := store.OrgGetByIDStore(ctx, id)
	if err != nil {
		return nil, err
	}

	// begin transaction
	tx, err := settings.BeginTx(ctx)
	if err != nil {
		return nil, err
	}
	defer settings.RollbackTx(ctx, tx)

	//Validation
	if req.Name != "" {
		org.Name = req.Name
	}

	if req.Website != "" {
		org.Website = req.Website
	}

	if req.Sector != "" {
		org.Sector = req.Sector
	}

	if req.Status != "" {
		org.Status = req.Status
	}

	err = store.UpdateOrgStore(ctx, tx, *org)
	if err != nil {
		return nil, err
	}

	if err := settings.CommitTx(ctx, tx); err != nil {
		return nil, err
	}

	return org, nil
}
