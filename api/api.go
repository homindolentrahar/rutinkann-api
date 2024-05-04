package api

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"com.homindolentrahar.rutinkann-api/controller"
	"com.homindolentrahar.rutinkann-api/helper"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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

func CreateApiService(authController controller.AuthController, routineController controller.RoutineController, logController controller.LogController) *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Timeout(60 * time.Second))
	router.Use(middleware.SetHeader("Content-Type", "application/json"))

	router.Route("/api/v1", func(r chi.Router) {
		r.Post("/sign-in", authController.SignIn)
		r.Post("/register", authController.Register)

		r.Group(func(r chi.Router) {
			r.Use(JWTAuthMiddleware())

			r.Route("/routines", func(r chi.Router) {
				r.Get("/", routineController.FindAll)
				r.Get("/{id}", routineController.FindById)
				r.Post("/", routineController.Create)
				r.Put("/{id}", routineController.Update)
				r.Delete("/{id}", routineController.Delete)
			})

			r.Route("/logs", func(r chi.Router) {
				r.Get("/", logController.FindAll)
				r.Get("/{id}", logController.FindById)
				r.Post("/", logController.Create)
				r.Put("/{id}", logController.Update)
				r.Delete("/{id}", logController.Delete)
			})

		})
	})

	// err := http.ListenAndServe(address, router)
	// if err != nil {
	// 	return err
	// }

	return router
}
