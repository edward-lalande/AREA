package utils

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func OpenDB(r *gin.Context) *pgx.Conn {
	connConfig, errConfig := pgx.ParseConfig(GetEnvKey("DB"))
	if errConfig != nil {
		return nil
	}

	db, dbOpenError := pgx.ConnectConfig(context.Background(), connConfig)
	if dbOpenError != nil {
		return nil
	}

	return db
}
