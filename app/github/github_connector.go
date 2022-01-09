package github

import (
	"context"

	"github.com/google/go-github/v41/github"
	"github.com/guiln/boilerplate-cli/src/models"
	"golang.org/x/oauth2"
)

const (
	folderContentType            string = "dir"
	fileContentType              string = "file"
	boilerplateIndicatorFileName string = ".boilerplate" //if folder contains this file in the root means that it is a boilerplate container folder
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

func (gc *GithubConnector) GetTemplateRepo() (*models.BoilerplateRepo, *models.BoilerplateError) {
	rootFolder, err := gc.traverseDirectory("/")
	if err != nil {
		return nil, err
	}

	boilerplateRepo := models.NewBoilerplateRepo(rootFolder)

	return boilerplateRepo, nil
}

func (gc *GithubConnector) Fetch(path string) *models.BoilerplateError {

	return nil
}

func (gc *GithubConnector) traverseDirectory(dirName string) (*models.BoilerplateFolder, *models.BoilerplateError) {
	_, dirContent, _, err := gc.client.Repositories.GetContents(gc.ctx, gc.options.GitBoilerplateRepositoryOwner, gc.options.GitBoilerplateRepository, dirName, &github.RepositoryContentGetOptions{})
	if err != nil {
		return nil, models.CreateBoilerplateErrorFromError(err, "error occured when traversing repo on github")
	}

	currentFolder := models.NewBoilerplateFolder(dirName, false)
	var childFoldersPathToTraverse []string

	for _, content := range dirContent {
		contentType := content.GetType()
		if contentType == fileContentType {
			if content.GetName() == boilerplateIndicatorFileName {
				currentFolder.SetIsContainer(true)
				return currentFolder, nil
			}
		} else if contentType == folderContentType {
			childFoldersPathToTraverse = append(childFoldersPathToTraverse, content.GetPath())
		}
	}

	for _, childPath := range childFoldersPathToTraverse {
		childRepo, berror := gc.traverseDirectory(childPath)
		if err != nil {
			return nil, berror
		}
		currentFolder.AddChild(childRepo)
	}

	return currentFolder, nil
}

func (gc *GithubConnector) initialize() {
	gc.ctx = context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: gc.options.GitToken})
	tc := oauth2.NewClient(gc.ctx, ts)

	gc.client = github.NewClient(tc)
}
