package utils

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func OpenDB(r *gin.Context) *pgx.Conn {
	connConfig, errConfig := pgx.ParseConfig(GetEnvKey("DB"))
	if errConfig != nil {
		if r != nil {
			r.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur de configuration de connexion"})
		} else {
			log.Printf("Error parsing DB config: %v", errConfig)
		}
		return nil
	}

	db, dbOpenError := pgx.ConnectConfig(context.Background(), connConfig)
	if dbOpenError != nil {
		if r != nil {
			r.JSON(http.StatusInternalServerError, gin.H{"error": dbOpenError.Error()})
		} else {
			log.Printf("Error connecting to DB: %v", dbOpenError)
		}
		return nil
	}

	return db
}
