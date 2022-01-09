package adapters

import "github.com/guiln/boilerplate-cli/src/models"

type ExternalRepoConnector interface {
	GetTemplateRepo() (*models.BoilerplateRepo, *models.BoilerplateError)
}
