package ports

import "github.com/guiln/boilerplate-cli/domain/models"

type TemplateFetcher interface {
	FetchRepo(path, folderPath string) *models.BoilerplateError
	FetchDoc(path, localPath string) *models.BoilerplateError
}
