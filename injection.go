//go:build wireinject
// +build wireinject

package main

import (
	"com.homindolentrahar.rutinkann-api/api"
	"com.homindolentrahar.rutinkann-api/controller"
	"com.homindolentrahar.rutinkann-api/db"
	"com.homindolentrahar.rutinkann-api/repository"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
)

var (
	authSet = wire.NewSet(
		repository.NewAuthRepositoryImpl,
		wire.Bind(new(repository.AuthRepository), new(*repository.AuthRepositoryImpl)),
		controller.NewAuthControllerImpl,
		wire.Bind(new(controller.AuthController), new(*controller.AuthControllerImpl)),
	)
	logSet = wire.NewSet(
		repository.NewLogRepository,
		wire.Bind(new(repository.LogRepository), new(*repository.LogRepositoryImpl)),
		controller.NewLogController,
		wire.Bind(new(controller.LogController), new(*controller.LogControllerImpl)),
	)
	routineSet = wire.NewSet(
		repository.NewRoutineRepository,
		wire.Bind(new(repository.RoutineRepository), new(*repository.RoutineRepositoryImpl)),
		controller.NewRoutineController,
		wire.Bind(new(controller.RoutineController), new(*controller.RoutineControllerImpl)),
	)
)

func ProvideValidatorOptions() []validator.Option {
	return []validator.Option{}
}

func InitializeServer() *chi.Mux {
	wire.Build(
		ProvideValidatorOptions,
		validator.New,
		db.ConnectPostgresStorage,
		authSet,
		routineSet,
		logSet,
		api.CreateApiService,
	)

	return nil
}
