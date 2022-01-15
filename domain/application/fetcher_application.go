package application

import (
	"fmt"

	"github.com/guiln/boilerplate-cli/domain/adapters"
	"github.com/guiln/boilerplate-cli/domain/models"
)

type FetcherApplication struct {
	options *FetcherApplicationOptions
}

func NewFetcherApplication(options *FetcherApplicationOptions) *FetcherApplication {
	return &FetcherApplication{options: options}
}

func (fApp *FetcherApplication) SyncWithTemplateRepo() {
}

func (fApp *FetcherApplication) GetLocalRepo() (*models.BoilerplateRepo, *models.BoilerplateError) {
	repo, err := fApp.options.RepoHandler.ReadRepo()
	if err != nil {
		return nil, models.CreateBoilerplateErrorFromError(err, "error while reading boilerplate repo in list operation")
	}

	return repo, nil
}

func (fApp *FetcherApplication) Fetch(repoPath, folderPath string) *models.BoilerplateError {
	if folderPath == "" {
		folderPath = repoPath
	}
	if err := fApp.options.TemplateFetcher.FetchRepo(repoPath, folderPath); err != nil {
		return models.CreateBoilerplateErrorFromError(err, fmt.Sprintf("failed to fetch %s repo", repoPath))
	}

	return nil
}

func (fApp *FetcherApplication) Sync() *models.BoilerplateError {
	templateRepo, err := fApp.options.ExternalRepoConnector.GetTemplateRepo()
	if err != nil {
		return err
	}
	if err := fApp.options.RepoHandler.PersistRepo(templateRepo); err != nil {
		return err
	}

	return nil
}

type FetcherApplicationOptions struct {
	RepoHandler           adapters.RepoHandler
	ExternalRepoConnector adapters.ExternalRepoConnector
	TemplateFetcher       adapters.TemplateFetcher
}
