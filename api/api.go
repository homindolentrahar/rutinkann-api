package api

import (
	"database/sql"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"time"
)

type RestApiService interface {
	StartServer() error
}

type ChiApiService struct {
	Address  string
	Database *sql.DB
}

func NewChiApiService(address string, database *sql.DB) *ChiApiService {
	return &ChiApiService{
		Address:  address,
		Database: database,
	}
}

func (chiApiService *ChiApiService) StartServer() error {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Timeout(60 * time.Second))

	ActivityRoute(router, chiApiService.Database)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}

	return nil
}
