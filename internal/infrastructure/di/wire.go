//go:build wireinject
// +build wireinject

package di

//go:generate wire

import (
	"github.com/3110Y/idsaver/internal/application/mapping"
	"github.com/3110Y/idsaver/internal/application/service"
	"github.com/3110Y/idsaver/internal/infrastructure/database"
	"github.com/3110Y/idsaver/internal/infrastructure/repository"
	"github.com/3110Y/idsaver/internal/infrastructure/router"
	httpcontroller "github.com/3110Y/idsaver/internal/presentation/http"
	"github.com/go-chi/chi/v5"
	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
)

type DI struct {
	DB           *sqlx.DB
	Router       *chi.Mux
	IdRepository *repository.IdRepository
	IdMapper     *mapping.IdMapper
	IdService    *service.IdService
	IdController *httpcontroller.IdController
}

func NewDI(
	DB *sqlx.DB,
	r *chi.Mux,
	IdRepository *repository.IdRepository,
	IdMapper *mapping.IdMapper,
	IdService *service.IdService,
	IdController *httpcontroller.IdController,
) *DI {
	return &DI{
		DB:           DB,
		Router:       r,
		IdRepository: IdRepository,
		IdMapper:     IdMapper,
		IdService:    IdService,
		IdController: IdController,
	}
}

func InitializeDI() (*DI, error) {
	wire.Build(
		NewDI,
		database.NewConnect,
		router.NewRouter,
		repository.NewIdRepository,
		mapping.NewIdMapper,
		service.NewIdService,
		httpcontroller.NewIdController,
	)
	return &DI{}, nil
}
