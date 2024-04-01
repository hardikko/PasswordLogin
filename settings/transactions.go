package settings

import (
	"context"
	"learngo/utils/faulterr"
	"log"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

var (
	DBClient *pgxpool.Pool // this is variable because we will assign new value once db is connected
)

func BeginTx(ctx context.Context) (pgx.Tx, *faulterr.FaultErr) {
	errMsg := "error when trying to begin"
	tx, err := DBClient.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, faulterr.NewPostgresError(err, errMsg)
	}
	log.Println("Beign Transaction")

	return tx, nil

}

func CommitTx(ctx context.Context, tx pgx.Tx) *faulterr.FaultErr {
	errMsg := "error when trying to get users"
	err := tx.Commit(ctx)
	if err != nil {
		return faulterr.NewPostgresError(err, errMsg)
	}
	log.Println("Commit Transaction")

	return nil
}

func RollbackTx(ctx context.Context, tx pgx.Tx) error {
	err := tx.Rollback(ctx)
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	log.Println("Rollback Transaction")

	return nil
}
