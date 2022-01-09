package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "lists templates contained locally, to sync with remote repo use sync command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("following templates are available:")
		templateRepo, err := fetcherApplication.GetLocalRepo()
		if err != nil {
			panic(err)
		}
		strRepoRepresentation, err := templateRepo.RootBoilerplateFolder.String()
		if err != nil {
			panic(err)
		}
		fmt.Println(strRepoRepresentation)
		fmt.Println("\n\nto use following templates use fetch command with desired template path separated by / like:")
		fmt.Println("tfetch node/api/express-restful")
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
