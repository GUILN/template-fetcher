package cmd

import (
	"fmt"

	"github.com/guiln/boilerplate-cli/domain/models"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var (
	repoTemplatePath string
	docTemplatePath  string
	localPathName    string
)

var fetchCmd = &cobra.Command{
	Use:   command_name_fetch,
	Short: "fetches template",
	Run: func(cmd *cobra.Command, args []string) {
		if (repoTemplatePath == "" && docTemplatePath == "") || (repoTemplatePath != "" && docTemplatePath != "") {
			promptOptions()
			return
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

// promptOptions
// this function is under test
func promptOptions() {
	templateRepo, err := fetcherApplication.GetLocalRepo()
	if err != nil {
		panic(err)
	}

	rootTemplateFolder := templateRepo.RootBoilerplateFolder
	index, result, _ := promptSelectOptionsForFolder(rootTemplateFolder.Path, rootTemplateFolder.ChildBoilerplateFolders)

	fmt.Printf("index: %d | path %s", index, result)
}

func promptSelectOptionsForFolder(rootFolderName string, folders []*models.BoilerplateFolder) (int, string, error) {
	options := getPromptOptionsFromBoilerplateFolders(folders)

	prompt := promptui.Select{
		Label: rootFolderName,
		Items: options,
	}

	index, result, errPrompt := prompt.Run()
	if errPrompt != nil {
		return -1, "", errPrompt
	}

	return index, result, nil
}

func getPromptOptionsFromBoilerplateFolders(folders []*models.BoilerplateFolder) []string {
	var boilerplateOptions []string
	for _, child := range folders {
		boilerplateOptions = append(boilerplateOptions, child.Path)
	}
	return boilerplateOptions
}

func init() {
	fetchCmd.PersistentFlags().StringVar(&repoTemplatePath, arg_name_repo, "", "template=[path/to/your/template_repo]")
	fetchCmd.PersistentFlags().StringVar(&docTemplatePath, arg_name_doc, "", "doc=[path/to/your/template_doc]")
	fetchCmd.PersistentFlags().StringVar(&localPathName, arg_name_name, "", "name=[path to dump template]")
	rootCmd.AddCommand(fetchCmd)
}
