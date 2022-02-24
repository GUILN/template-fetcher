package cmd

import (
	"github.com/guiln/boilerplate-cli/domain/models"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var (
	repoTemplatePath string
	docTemplatePath  string
	localPathName    string
)

type templateType string

const (
	templateRepoType templateType = "Repo"
	templateDocType  templateType = "Doc"
)

var fetchCmd = &cobra.Command{
	Use:   command_name_fetch,
	Short: "fetches template",
	Run: func(cmd *cobra.Command, args []string) {
		// If none or both repo and doc template path are provided
		// Prompts the options for a interactive select
		if (repoTemplatePath == "" && docTemplatePath == "") || (repoTemplatePath != "" && docTemplatePath != "") {
			templatePath, tType, err := promptOptions()
			if err != nil {
				panic(err)
			}
			if tType == templateRepoType {
				repoTemplatePath = templatePath
			} else {
				docTemplatePath = templatePath
			}
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
// Return: (templatePath: str, templateType: templateType, error: error)
func promptOptions() (string, templateType, error) {
	templateRepo, err := fetcherApplication.GetLocalRepo()
	if err != nil {
		panic(err)
	}

	currentTemplateFolder := templateRepo.RootBoilerplateFolder
	var (
		finalPath string       = ""
		pathType  templateType = ""
	)
	for (currentTemplateFolder.ChildTemplateDocuments != nil || currentTemplateFolder.ChildBoilerplateFolders != nil) && pathType == "" {
		var (
			index        int
			promptResult string
			err          error
		)

		if currentTemplateFolder.IsDocContainerFolder {
			_, promptResult, err = promptSelectOptionsForDocs(currentTemplateFolder.Path, currentTemplateFolder.ChildTemplateDocuments)
			pathType = templateDocType
		} else {
			index, promptResult, err = promptSelectOptionsForFolder(currentTemplateFolder.Path, currentTemplateFolder.ChildBoilerplateFolders)
			currentTemplateFolder = currentTemplateFolder.ChildBoilerplateFolders[index]
			if currentTemplateFolder.IsRepoContainerFolder {
				pathType = templateRepoType
			}
		}

		if err != nil {
			return "", "", err
		}
		finalPath = promptResult
	}

	return finalPath, pathType, nil
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
