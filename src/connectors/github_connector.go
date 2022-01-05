package connectors

import (
	"context"
	"fmt"

	"github.com/google/go-github/v41/github"
	"github.com/guiln/boilerplate-cli/src/models"
	"golang.org/x/oauth2"
)

type GithubConnector struct {
	client  *github.Client
	options *GithubConnectorOptions
	ctx     context.Context
}

type GithubConnectorOptions struct {
	GitToken                      string
	GitBoilerplateRepository      string
	GitBoilerplateRepositoryOwner string
}

func NewGithubConnector(options *GithubConnectorOptions) *GithubConnector {
	conn := &GithubConnector{options: options}
	conn.initialize()

	return conn
}

func (gc *GithubConnector) GetBoilerplates() (*models.BoilerplateRepo, *models.BoilerplateError) {
	_, dirContent, _, err := gc.client.Repositories.GetContents(gc.ctx, gc.options.GitBoilerplateRepositoryOwner, gc.options.GitBoilerplateRepository, "/", &github.RepositoryContentGetOptions{})
	if err != nil {
		return nil, &models.BoilerplateError{Message: "error occured during fetch repo on github", InnerError: err}
	}

	for _, dir := range dirContent {
		fmt.Printf("Dir Content: \n%v", dir)
	}

	return &models.BoilerplateRepo{}, nil
}

func (gc *GithubConnector) initialize() {
	gc.ctx = context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: gc.options.GitToken})
	tc := oauth2.NewClient(gc.ctx, ts)

	gc.client = github.NewClient(tc)
}
