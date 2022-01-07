package sync

type FetcherApplication struct {
	options *FetcherApplicationOptions
}

func NewFetcherApplication(options *FetcherApplicationOptions) *FetcherApplication {
	return &FetcherApplication{options: options}
}

func (fApp *FetcherApplication) SyncWithTemplateRepo() {
}

func (fApp *FetcherApplication) List() {
}

func (fApp *FetcherApplication) Fetch(repoPath ...string) {
}

type FetcherApplicationOptions struct {
	Repo      string
	RepoOwner string
	Token     string
}
