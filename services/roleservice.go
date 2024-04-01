package services

import (
	"context"
	"fmt"
	"learngo/models"
	"learngo/settings"
	"learngo/store"
	"learngo/utils/faulterr"
)

func RoleService(ctx context.Context, id int64) (*models.Role, *faulterr.FaultErr) {
	obj, err := store.RoleGetByIDStore(ctx, id)
	if err != nil {
		return nil, err
	}

	return obj, nil
}

func RoleCreateService(ctx context.Context, req models.Role) (*models.Role, *faulterr.FaultErr) {
	//Get all Role list
	rol, err := store.RoleListStore(ctx)

	if err != nil {
		return nil, err
	}
	code := fmt.Sprintf("ROLE%03d", len(rol)+1)
	req.Code = code

	// connect to dbstore
	tx, err := settings.BeginTx(ctx)
	if err != nil {
		return nil, err
	}
	defer settings.RollbackTx(ctx, tx)

	obj, err := store.RoleInsertStore(ctx, tx, req)
	if err != nil {
		return nil, err
	}

	if err := settings.CommitTx(ctx, tx); err != nil {
		return nil, err
	}

	return obj, nil
}

func RoleListService(ctx context.Context) ([]models.Role, *faulterr.FaultErr) {
	result, err := store.RoleListStore(ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func RoleUpdateService(ctx context.Context, id int64, req models.Role) (*models.Role, *faulterr.FaultErr) {
	// get role by id from db
	role, err := store.RoleGetByIDStore(ctx, id)
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
		role.Name = req.Name
	}

	if req.Status != "" {
		role.Status = req.Status
	}

	err = store.UpdateRoleStore(ctx, tx, *role)
	if err != nil {
		return nil, err
	}

	if err := settings.CommitTx(ctx, tx); err != nil {
		return nil, err
	}

	return role, nil
}
