package adapters

import "github.com/guiln/boilerplate-cli/src/models"

type RepoHandler interface {
	PersistRepo(*models.BoilerplateRepo) *models.BoilerplateError
	ReadRepo() (*models.BoilerplateRepo, *models.BoilerplateError)
}
