package ports

import "github.com/guiln/boilerplate-cli/domain/models"

type RepoHandler interface {
	PersistRepo(*models.BoilerplateRepo) *models.BoilerplateError
	ReadRepo() (*models.BoilerplateRepo, *models.BoilerplateError)
}
