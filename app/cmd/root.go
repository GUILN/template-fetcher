package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/guiln/boilerplate-cli/app/configuration"
	"github.com/guiln/boilerplate-cli/app/github"
	"github.com/guiln/boilerplate-cli/app/repo"
	"github.com/guiln/boilerplate-cli/src/application"
	"github.com/spf13/cobra"
)

const (
	hardcodedSecret string = "abc&1*~#^2^#s0^=)^^7%b34"
)

var (
	cfg                *configuration.Config
	fetcherApplication *application.FetcherApplication
)

var rootCmd = &cobra.Command{
	Use:   "cmd",
	Short: "A template fetcher to help initiliazing projects with boilerplate code / references / packages / libraries",
	Long: `Ideal tool for standarized templates that can be stored and fetched from github, and other common repos.
	Helps the team building microservices architecture to have templates in one place.
	Helps your personal projects that you are tired of copy and paste from somewhere else.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig, initApplication)

	rootCmd.CompletionOptions.DisableDefaultCmd = true
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	cfgDir := filepath.Join(userHomeDir, ".template.fetcher")
	cfg = configuration.NewConfig(cfgDir, hardcodedSecret)
	if err := cfg.LoadConfig(); err != nil {
		fmt.Println("error while trying to load configuration, please make sure you have set the configuration with config command")
		fmt.Printf("%v", err)
	}
}

func initApplication() {
	tkn, err := cfg.GetToken()
	if err != nil {
		panic(err)
	}
	ghbConnector := github.NewGithubConnector(&github.GithubConnectorOptions{
		GitToken:                      tkn,
		GitBoilerplateRepository:      cfg.Repo,
		GitBoilerplateRepositoryOwner: cfg.RepoOwner,
	})
	fetcherApplication = application.NewFetcherApplication(&application.FetcherApplicationOptions{
		RepoHandler:           repo.NewTemplateJsonRepo(cfg.GetTemplatesFilePath()),
		ExternalRepoConnector: ghbConnector,
		TemplateFetcher:       ghbConnector,
	})
}
