package adapters

import "github.com/guiln/boilerplate-cli/domain/models"

type ExternalRepoConnector interface {
	GetTemplateRepo() (*models.BoilerplateRepo, *models.BoilerplateError)
}
