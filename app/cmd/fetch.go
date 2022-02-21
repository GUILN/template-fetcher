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
			selectedOption, _ := promptOptions()
			fmt.Print(selectedOption)
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
// this function returns the path mounted from prompt
func promptOptions() (string, error) {
	templateRepo, err := fetcherApplication.GetLocalRepo()
	if err != nil {
		panic(err)
	}

	currentTemplateFolder := templateRepo.RootBoilerplateFolder
	var (
		finalPath string = ""
		pathType  string = ""
	)
	for (currentTemplateFolder.ChildTemplateDocuments != nil || currentTemplateFolder.ChildBoilerplateFolders != nil) && pathType == "" {
		var (
			index        int
			promptResult string
			err          error
		)

		if currentTemplateFolder.IsDocContainerFolder {
			_, promptResult, err = promptSelectOptionsForDocs(currentTemplateFolder.Path, currentTemplateFolder.ChildTemplateDocuments)
			pathType = "Document"
		} else {
			index, promptResult, err = promptSelectOptionsForFolder(currentTemplateFolder.Path, currentTemplateFolder.ChildBoilerplateFolders)
			currentTemplateFolder = currentTemplateFolder.ChildBoilerplateFolders[index]
			if currentTemplateFolder.IsRepoContainerFolder {
				pathType = "Repo"
			}
		}

		if err != nil {
			return "", err
		}
		finalPath = promptResult
	}

	return finalPath, nil
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

func promptSelectOptionsForDocs(rootFolderName string, docs []*models.TemplateDocument) (int, string, error) {
	options := getPromptOptionsFromBoilerplateDocs(docs)

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

func getPromptOptionsFromBoilerplateDocs(docs []*models.TemplateDocument) []string {
	var boilerplateOptions []string
	for _, child := range docs {
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
