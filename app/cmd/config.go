package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var (
	repoName   string
	repoOwner  string
	gitToken   string
	listConfig bool
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "use config command to configure tfetcher paramaters",
	Run: func(cmd *cobra.Command, args []string) {
		if listConfig {
			fmt.Println(cfg.String())
		} else {
			updateConfig()
		}
	},
}

func updateConfig() {
	updatedParametersBuilder := strings.Builder{}

	if repoName != "" {
		cfg.Repo = repoName
		updatedParametersBuilder.WriteString("repo name updated!\n")
	}
	if repoOwner != "" {
		cfg.RepoOwner = repoOwner
		updatedParametersBuilder.WriteString("repo owner updated!\n")
	}
	if gitToken != "" {
		if cfgError := cfg.SetToken(gitToken); cfgError != nil {
			panic(cfgError)
		}
		updatedParametersBuilder.WriteString("repo token updated!\n")
	}

	if cfgError := cfg.PersistConfig(); cfgError != nil {
		panic(cfgError)
	}

	if updatedParameters := updatedParametersBuilder.String(); updatedParameters != "" {
		fmt.Println("updated parameters: ")
		fmt.Println(updatedParameters)
	}
}

func init() {
	configCmd.PersistentFlags().StringVar(&repoName, "repo-name", "", "repo-name=[your repository name]")
	configCmd.PersistentFlags().StringVar(&repoOwner, "repo-owner", "", "repo-owner=[repository's owner]")
	configCmd.PersistentFlags().StringVar(&gitToken, "token", "", "token=[your secret token] (stored encrypted)")
	configCmd.PersistentFlags().BoolVar(&listConfig, "list", false, "list")
	rootCmd.AddCommand(configCmd)
}
