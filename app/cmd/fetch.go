package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	repoTemplatePath string
	docTemplatePath  string
	localPathName    string
)

var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "fetches template",
	Run: func(cmd *cobra.Command, args []string) {
		if (repoTemplatePath == "" && docTemplatePath == "") || (repoTemplatePath != "" && docTemplatePath != "") {
			fmt.Println("exactly only one flag from repo or doc should be specified")
			os.Exit(1)
		}

		if repoTemplatePath != "" {
			if err := fetcherApplication.Fetch(repoTemplatePath, localPathName); err != nil {
				panic(err)
			}
		} else {
			if err := fetcherApplication.FetchDoc(docTemplatePath, localPathName); err != nil {
				panic(err)
			}
		}

	},
}

func init() {
	fetchCmd.PersistentFlags().StringVar(&repoTemplatePath, "repo", "", "template=[path/to/your/template_repo]")
	fetchCmd.PersistentFlags().StringVar(&docTemplatePath, "doc", "", "doc=[path/to/your/template_doc]")
	fetchCmd.PersistentFlags().StringVar(&localPathName, "name", "", "name=[path to dump template]")
	rootCmd.AddCommand(fetchCmd)
}
