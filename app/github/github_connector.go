package github

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/google/go-github/v41/github"
	"github.com/guiln/boilerplate-cli/domain/models"
	"github.com/guiln/boilerplate-cli/helpers"
	"golang.org/x/oauth2"
)

const (
	folderContentType            string = "dir"
	fileContentType              string = "file"
	boilerplateIndicatorFileName string = ".boilerplate" //if folder contains this file in the root means that it is a boilerplate container folder
	docDirIndicatorFileName      string = ".docdir"      //if folder contains this file in the root means that it is a doc template container folder
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

func (gc *GithubConnector) Fetch(path, folderPath string) *models.BoilerplateError {
	_, dirContent, _, err := gc.client.Repositories.GetContents(gc.ctx, gc.options.GitBoilerplateRepositoryOwner, gc.options.GitBoilerplateRepository, path, &github.RepositoryContentGetOptions{})
	if err != nil {
		return models.CreateBoilerplateErrorFromError(err, "error occured when traversing repo on github")
	}

	fmt.Printf("creating dir %s...\n", folderPath)

	if err := os.MkdirAll(folderPath, 0755); err != nil {
		return models.CreateBoilerplateErrorFromError(err, "error occured when trying to create repo's folder locally")
	}
	var subDirsPath []string

	for _, content := range dirContent {
		contentType := content.GetType()
		if contentType == fileContentType {
			fileName := content.GetName()
			downloadUrl := content.GetDownloadURL()
			fullFileName := filepath.Join(folderPath, fileName)

			fmt.Printf("downloading %s at %s...\n", fullFileName, downloadUrl)
			if err := gc.downloadFile(content, fullFileName); err != nil {
				return err
			}
			// Downoads content
		} else if contentType == folderContentType {
			subDirsPath = append(subDirsPath, content.GetPath())
		}
	}

	for _, subDirPath := range subDirsPath {
		subDirFolderPath := filepath.Join(folderPath, filepath.Base(subDirPath))
		if err := gc.Fetch(subDirPath, subDirFolderPath); err != nil {
			return err
		}
	}

	return nil
}

func (gc *GithubConnector) traverseDirectory(dirName string) (*models.BoilerplateFolder, *models.BoilerplateError) {
	_, dirContent, _, err := gc.client.Repositories.GetContents(gc.ctx, gc.options.GitBoilerplateRepositoryOwner, gc.options.GitBoilerplateRepository, dirName, &github.RepositoryContentGetOptions{})
	if err != nil {
		return nil, models.CreateBoilerplateErrorFromError(err, "error occured when traversing repo on github")
	}

	currentFolder := models.NewBoilerplateFolder(dirName, false, false)
	var childFoldersPathToTraverse []string

	for _, content := range dirContent {
		contentType := content.GetType()
		if contentType == fileContentType {
			if content.GetName() == boilerplateIndicatorFileName {
				currentFolder.SetIsRepoContainer(true)
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

func (gc *GithubConnector) downloadFile(repoContent *github.RepositoryContent, downloadPath string) *models.BoilerplateError {
	repoPath := repoContent.GetPath()
	readCloser, _, err := gc.client.Repositories.DownloadContents(gc.ctx, gc.options.GitBoilerplateRepositoryOwner, gc.options.GitBoilerplateRepository, repoPath, &github.RepositoryContentGetOptions{})
	if err != nil {
		return models.CreateBoilerplateErrorFromError(err, fmt.Sprintf("error occured when downloading file %s", repoPath))
	}
	defer readCloser.Close()

	responseBytes, err := ioutil.ReadAll(readCloser)
	if err != nil {
		return models.CreateBoilerplateErrorFromError(err, fmt.Sprintf("error occured when downloading file %s", repoPath))
	}

	fmt.Printf("writing file: %s\n", downloadPath)
	if err := helpers.CreateFile(responseBytes, downloadPath); err != nil {
		fmt.Print(err)
		return models.CreateBoilerplateErrorFromError(err, fmt.Sprintf("error occured when trying to create file %s", downloadPath))
	}

	return nil
}

func (gc *GithubConnector) initialize() {
	gc.ctx = context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: gc.options.GitToken})
	tc := oauth2.NewClient(gc.ctx, ts)

	gc.client = github.NewClient(tc)
}
