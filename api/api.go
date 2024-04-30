package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"com.homindolentrahar.rutinkann-api/controller"
	"com.homindolentrahar.rutinkann-api/helper"
	"com.homindolentrahar.rutinkann-api/repository"
	"github.com/go-playground/validator/v10"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gorm.io/gorm"
)

func JWTAuthMiddleware() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			encoder := json.NewEncoder(w)
			if r.Header.Get("Authorization") == "" || !strings.Contains(r.Header.Get("Authorization"), "Bearer") {
				w.WriteHeader(http.StatusUnauthorized)
				response := map[string]interface{}{
					"status":  http.StatusUnauthorized,
					"message": "Unauthorized request",
				}
				encodeErr := encoder.Encode(response)
				if encodeErr != nil {
					return
				}

				return
			}

			secretKey := os.Getenv("APP_SECRET_KEY")
			authorization := r.Header.Get("Authorization")
			tokenString := strings.Replace(authorization, "Bearer ", "", -1)
			token, tokenErr := helper.VerifyToken(secretKey, tokenString)

			if tokenErr != nil {
				w.WriteHeader(http.StatusUnauthorized)
				response := map[string]interface{}{
					"status":  http.StatusUnauthorized,
					"message": tokenErr.Error(),
				}
				encodeErr := encoder.Encode(response)
				if encodeErr != nil {
					return
				}

				return
			}

			log.Println(token.Claims)

			next.ServeHTTP(w, r)
		}

		return http.HandlerFunc(fn)
	}
}

type RestApiService interface {
	StartServer() error
}

type ChiApiService struct {
	Address  string
	Database *gorm.DB
	Validate *validator.Validate
}

func NewChiApiService(address string, database *gorm.DB, validate *validator.Validate) *ChiApiService {
	return &ChiApiService{
		Address:  address,
		Database: database,
		Validate: validate,
	}
}

func (chiApiService *ChiApiService) StartServer() error {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Timeout(60 * time.Second))
	router.Use(middleware.SetHeader("Content-Type", "application/json"))

	router.Route("/api/v1", func(r chi.Router) {
		authRepository := repository.NewAuthRepositoryImpl(chiApiService.Validate)
		authController := controller.NewAuthControllerImpl(authRepository, chiApiService.Database)

		r.Post("/sign-in", authController.SignIn)
		r.Post("/register", authController.Register)

		r.Group(func(r chi.Router) {
			r.Use(JWTAuthMiddleware())

			r.Route("/activities", func(r chi.Router) {
				activityRepository := repository.NewActivityRepository()
				activityController := controller.NewActivityController(activityRepository, chiApiService.Database)

				r.Get("/", activityController.FindAll)
				r.Get("/{id}", activityController.FindById)
				r.Post("/", activityController.Create)
				r.Put("/{id}", activityController.Update)
				r.Delete("/{id}", activityController.Delete)
			})

			r.Route("/logs", func(r chi.Router) {
				logRepository := repository.NewLogRepository()
				logController := controller.NewLogController(logRepository, chiApiService.Database)

				r.Get("/", logController.FindAll)
				r.Get("/{id}", logController.FindById)
				r.Post("/", logController.Create)
				r.Put("/{id}", logController.Update)
				r.Delete("/{id}", logController.Delete)
			})
		})
	})

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}

	return nil
}
