package services

import (
	"context"
	"fmt"
	"learngo/helpers"
	"learngo/models"
	"learngo/settings"
	"learngo/store"
	"learngo/utils/faulterr"
)

func UserListService(ctx context.Context) ([]models.User, *faulterr.FaultErr) {
	result, err := store.UserListStore(ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func UserService(ctx context.Context, id int64) (*models.User, *faulterr.FaultErr) {
	obj, err := store.UserGetByIDStore(ctx, id)
	if err != nil {
		return nil, err
	}

	return obj, nil
}

func UserCreateService(ctx context.Context, req models.User) (*models.User, *faulterr.FaultErr) {
	//Get all Organization list
	users, err := store.UserListStore(ctx)

	if err != nil {
		return nil, err
	}
	code := fmt.Sprintf("USER%03d", len(users)+1)
	req.Code = code

	// connect to dbstore
	tx, err := settings.BeginTx(ctx)
	if err != nil {
		return nil, err
	}
	defer settings.RollbackTx(ctx, tx)

	//Create Empty variable
	req.PasswordHash = helpers.GetMd5(req.PasswordHash)

	obj, err := store.UserInsertStore(ctx, tx, req)
	if err != nil {
		return nil, err
	}

	if err := settings.CommitTx(ctx, tx); err != nil {
		return nil, err
	}

	return obj, nil
}

func UserUpdateService(ctx context.Context, id int64, req models.User) (*models.User, *faulterr.FaultErr) {
	// get user by id from db
	user, err := store.UserGetByIDStore(ctx, id)
	if err != nil {
		return nil, err
	}

	// begin transaction
	tx, err := settings.BeginTx(ctx)
	if err != nil {
		return nil, err
	}
	defer settings.RollbackTx(ctx, tx)
	user.FirstName = req.FirstName
	user.LastName = req.LastName
	user.Email = req.Email
	user.Phone = req.Phone
	user.Status = req.Status

	err = store.UpdateUserStore(ctx, tx, *user)
	if err != nil {
		return nil, err
	}

	if err := settings.CommitTx(ctx, tx); err != nil {
		return nil, err
	}

	return user, nil
}
