package sync

import "github.com/guiln/boilerplate-cli/src/configuration"

type FetcherApplication struct {
	config *configuration.Config
}

func NewFetcherApplication(config *configuration.Config) *FetcherApplication {
	return &FetcherApplication{config: config}
}

func (fApp *FetcherApplication) SyncWithTemplateRepo() {
}

func (fApp *FetcherApplication) List() {
}

func (fApp *FetcherApplication) Fetch(repoPath ...string) {
}
