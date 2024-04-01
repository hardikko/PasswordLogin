package services

import (
	"context"

	"learngo/models"
	"learngo/settings"
	"learngo/store"
	"learngo/utils/faulterr"
)

func LeadService(ctx context.Context, id int64) (*models.Lead, *faulterr.FaultErr) {
	obj, err := store.LeadGetByIDStore(ctx, id)
	if err != nil {
		return nil, err
	}

	return obj, nil
}

func LeadCreateService(ctx context.Context, req models.Lead) (*models.Lead, *faulterr.FaultErr) {
	// connect to dbstore
	tx, err := settings.BeginTx(ctx)
	if err != nil {
		return nil, err
	}
	defer settings.RollbackTx(ctx, tx)

	obj, err := store.LeadInsertStore(ctx, tx, req)
	if err != nil {
		return nil, err
	}

	if err := settings.CommitTx(ctx, tx); err != nil {
		return nil, err
	}

	return obj, nil
}

func LeadListService(ctx context.Context) ([]models.Lead, *faulterr.FaultErr) {
	result, err := store.LeadListStore(ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func LeadUpdateService(ctx context.Context, id int64, req models.Lead) (*models.Lead, *faulterr.FaultErr) {
	// get lead by id from db
	org, err := store.LeadGetByIDStore(ctx, id)
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
	if req.FirstName != "" {
		org.FirstName = req.FirstName
	}

	if req.LastName != "" {
		org.LastName = req.LastName
	}

	if req.Address != "" {
		org.Address = req.Address
	}

	if req.Phone != "" {
		org.Phone = req.Phone
	}

	if req.Email != "" {
		org.Email = req.Email
	}

	if req.Occupation != "" {
		org.Occupation = req.Occupation
	}

	if req.Status != "" {
		org.Status = req.Status
	}
	err = store.UpdateLeadStore(ctx, tx, *org)
	if err != nil {
		return nil, err
	}

	if err := settings.CommitTx(ctx, tx); err != nil {
		return nil, err
	}

	return org, nil
}
