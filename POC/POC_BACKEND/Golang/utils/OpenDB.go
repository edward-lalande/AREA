package utils

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func OpenDB(r *gin.Context) *pgx.Conn {
	connConfig, errConfig := pgx.ParseConfig("postgresql://root:password@localhost:5432/user")

	if errConfig != nil {
		r.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur de configuration de connexion"})
		return nil
	}
	db, dbOpenError := pgx.ConnectConfig(context.Background(), connConfig)
	if dbOpenError != nil {
		r.JSON(http.StatusInternalServerError, gin.H{"error": dbOpenError.Error()})
		return nil
	}

	return db
}
