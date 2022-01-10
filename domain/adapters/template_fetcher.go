package adapters

import "github.com/guiln/boilerplate-cli/domain/models"

type TemplateFetcher interface {
	Fetch(path string) *models.BoilerplateError
}
