package cmd

import (
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   command_name_list,
	Short: "lists templates contained locally, to sync with remote repo use sync command",
	Run: func(cmd *cobra.Command, args []string) {
		templateRepo, err := fetcherApplication.GetLocalRepo()
		if err != nil {
			panic(err)
		}
		strRepoRepresentation, err := templateRepo.RootBoilerplateFolder.String()
		if err != nil {
			panic(err)
		}
		printTemplates(strRepoRepresentation)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
