package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var templatePath string

var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "fetches template",
	Run: func(cmd *cobra.Command, args []string) {
		if templatePath == "" {
			panic(fmt.Errorf("path argument needs to be specified"))
		}

		if err := fetcherApplication.Fetch(templatePath); err != nil {
			panic(err)
		}
	},
}

func init() {
	fetchCmd.PersistentFlags().StringVar(&templatePath, "path", "", "path=[path/to/your/repo]")
	rootCmd.AddCommand(fetchCmd)
}
