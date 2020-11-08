package repository

import (
	core "github.com/skinnyguy/spectra-services/core/proto"
	repoQuery "github.com/skinnyguy/spectra-services/microservices/repository/query"

	"github.com/skinnyguy/spectra-services/core/utils"
)

type Repository interface {
	AddAccountQuery(data *core.AddAccountRequest) (string, error)
}

type AbstractRepository struct {
	Repository
}

func NewRepository() *AbstractRepository {
	newRepo := &AbstractRepository{}

	if utils.GetDatasourceInfo() == utils.Postgres {
		newRepo.Repository = &repoQuery.Repository{
			Connection: utils.GetPostgresHandler(),
		}
	} else {
		newRepo.Repository = nil
	}

	return newRepo
}