package github

import "testing"

func Test_GithubConnector_GetBoilerplates_GetsBoilerplatesSuccessfully(t *testing.T) {
	githubConnector := NewGithubConnector(&GithubConnectorOptions{
		GitToken:                      "",
		GitBoilerplateRepository:      "boilerplates",
		GitBoilerplateRepositoryOwner: "guiln",
	})

	githubConnector.GetTemplateRepo()
}
