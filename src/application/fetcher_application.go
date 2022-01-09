package application

import (
	"github.com/guiln/boilerplate-cli/src/adapters"
	"github.com/guiln/boilerplate-cli/src/models"
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

func (fApp *FetcherApplication) Fetch(repoPath ...string) {
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
}
