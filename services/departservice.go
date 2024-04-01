package services

import (
	"context"
	"fmt"
	"learngo/models"
	"learngo/settings"
	"learngo/store"
	"learngo/utils/faulterr"
)

func DepartService(ctx context.Context, id int64) (*models.Department, *faulterr.FaultErr) {
	obj, err := store.DepartGetByIDStore(ctx, id)
	if err != nil {
		return nil, err
	}

	return obj, nil
}

func DepartCreateService(ctx context.Context, req models.Department) (*models.Department, *faulterr.FaultErr) {
	//Get all Department list
	departs, err := store.DepartListStore(ctx)

	if err != nil {
		return nil, err
	}
	code := fmt.Sprintf("DEPART%03d", len(departs)+1)
	req.Code = code

	// connect to dbstore
	tx, err := settings.BeginTx(ctx)
	if err != nil {
		return nil, err
	}
	defer settings.RollbackTx(ctx, tx)

	obj, err := store.DepartInsertStore(ctx, tx, req)
	if err != nil {
		return nil, err
	}

	if err := settings.CommitTx(ctx, tx); err != nil {
		return nil, err
	}

	return obj, nil
}

func DepartListService(ctx context.Context) ([]models.Department, *faulterr.FaultErr) {
	result, err := store.DepartListStore(ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func DepartUpdateService(ctx context.Context, id int64, req models.Department) (*models.Department, *faulterr.FaultErr) {
	// get department by id from db
	depart, err := store.DepartGetByIDStore(ctx, id)
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
		depart.Name = req.Name
	}

	if req.Status != "" {
		depart.Status = req.Status
	}

	err = store.UpdateDepartStore(ctx, tx, *depart)
	if err != nil {
		return nil, err
	}

	if err := settings.CommitTx(ctx, tx); err != nil {
		return nil, err
	}

	return depart, nil
}
