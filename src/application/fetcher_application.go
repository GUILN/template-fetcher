package sync

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

func (fApp *FetcherApplication) List() (*models.BoilerplateRepo, *models.BoilerplateError) {
	repo, err := fApp.options.RepoHandler.ReadRepo()
	if err != nil {
		return nil, models.CreateBoilerplateErrorFromError(err, "error while reading boilerplate repo in list operation")
	}

	return repo, nil
}

func (fApp *FetcherApplication) Fetch(repoPath ...string) {
}

type FetcherApplicationOptions struct {
	Repo        string
	RepoOwner   string
	Token       string
	RepoHandler adapters.RepoHandler
}
