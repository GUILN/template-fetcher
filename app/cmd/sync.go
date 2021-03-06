package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// syncCmd represents the sync command
var syncCmd = &cobra.Command{
	Use:   command_name_sync,
	Short: "Sync local template repo with remote repo",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("starting sync...")
		if err := fetcherApplication.Sync(); err != nil {
			panic(err)
		}

		templateRepo, err := fetcherApplication.GetLocalRepo()
		if err != nil {
			panic(err)
		}

		strRepoRepresentation, err := templateRepo.RootBoilerplateFolder.String()
		if err != nil {
			panic(err)
		}

		fmt.Println("sync ended!")
		fmt.Println("local repo has been updated to:")
		printTemplates(strRepoRepresentation)
		printSyncedSuccessfully()
	},
}

func init() {
	rootCmd.AddCommand(syncCmd)
}
