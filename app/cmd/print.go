package cmd

import (
	"fmt"

	"github.com/fatih/color"
)

const separator = "----------------------------------------------------"

func printTemplates(templatesRepoRepresentation string) {
	color.Cyan("following templates are available:")
	color.Yellow(templatesRepoRepresentation)
	color.Cyan(separator)
	printFetchCommandFullExample()
}

func printFetchCommandFullExample() {
	color.Cyan(fetch_command_usage_text)
	color.Cyan(fetch_command_usage_text_repo)
	color.Cyan(fetch_command_usage_text_doc)
}

func printSyncedSuccessfully() {
	color.Cyan("sync-ed successfully with remote repo at %s!\n\n", cfg.Repo)
	fmt.Println("🍔🍟🍺🍺")
}
