package adapters

import "github.com/guiln/boilerplate-cli/src/models"

type RepoHandler interface {
	SaveRepo(*models.BoilerplateRepo) error
	ReadRepo() (*models.BoilerplateRepo, error)
}
