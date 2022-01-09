package adapters

import "github.com/guiln/boilerplate-cli/src/models"

type TemplateFetcher interface {
	Fetch(path string) *models.BoilerplateError
}
