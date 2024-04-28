package api

import (
	"com.homindolentrahar.rutinkann-api/controller"
	"com.homindolentrahar.rutinkann-api/repository"
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

func ActivityRoute(router *chi.Mux, db *gorm.DB) {
	activityRepository := repository.NewActivityRepository()
	activityController := controller.NewActivityController(activityRepository, db)

	router.Get("/api/v1/activities", activityController.FindAll)
	router.Get("/api/v1/activities/{id}", activityController.FindById)
	router.Post("/api/v1/activities", activityController.Create)
	router.Put("/api/v1/activities/{id}", activityController.Update)
	router.Delete("/api/v1/activities/{id}", activityController.Delete)
}

func LogRoute(router *chi.Mux, db *gorm.DB) {
	logRepository := repository.NewLogRepository()
	logController := controller.NewLogController(logRepository, db)

	router.Get("/api/v1/logs", logController.FindAll)
	router.Get("/api/v1/logs/{id}", logController.FindById)
	router.Post("/api/v1/logs", logController.Create)
	router.Put("/api/v1/logs/{id}", logController.Update)
	router.Delete("/api/v1/logs/{id}", logController.Delete)
}
