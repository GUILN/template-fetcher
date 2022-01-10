package adapters

import "github.com/guiln/boilerplate-cli/domain/models"

type TemplateFetcher interface {
	Fetch(path, folderPath string) *models.BoilerplateError
}
