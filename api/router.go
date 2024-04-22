package api

import (
	"com.homindolentrahar.rutinkann-api/controller"
	"com.homindolentrahar.rutinkann-api/data/web"
	"com.homindolentrahar.rutinkann-api/model"
	"com.homindolentrahar.rutinkann-api/repository"
	"database/sql"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func ActivityRoute(router *chi.Mux, db *sql.DB) {
	activityRepository := repository.NewActivityRepository()
	activityController := controller.NewActivityController(activityRepository, db)

	router.Get("/api/v1/activities", activityController.FindAll)
	router.Get("/api/v1/activities/{id}", activityController.FindById)
	router.Post("/api/v1/activities", activityController.Create)
	router.Put("/api/v1/activities/{id}", activityController.Update)
	router.Delete("/api/v1/activities/{id}", activityController.Delete)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		encoder := json.NewEncoder(w)

		encoder.Encode(
			web.BaseResponse[[]model.Activity]{
				Status:  http.StatusOK,
				Message: "Success",
				Data:    []model.Activity{},
			},
		)
	})
}
