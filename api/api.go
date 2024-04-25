package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gorm.io/gorm"
)

type RestApiService interface {
	StartServer() error
}

type ChiApiService struct {
	Address  string
	Database *gorm.DB
}

func NewChiApiService(address string, database *gorm.DB) *ChiApiService {
	return &ChiApiService{
		Address:  address,
		Database: database,
	}
}

func (chiApiService *ChiApiService) StartServer() error {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Timeout(60 * time.Second))
	router.Use(middleware.SetHeader("Content-Type", "application/json"))

	ActivityRoute(router, chiApiService.Database)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}

	return nil
}
