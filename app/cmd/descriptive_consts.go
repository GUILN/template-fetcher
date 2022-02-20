package cmd

import "fmt"

const (
	// Descrition of the commands
	command_name       string = "tfetch"
	command_name_fetch string = "fetch"
	command_name_list  string = "list"
	command_name_sync  string = "sync"

	// Description of the arguments
	arg_name_repo string = "repo"
	arg_name_doc  string = "doc"
	arg_name_name string = "name"
)

var (
	// Descriptive text
	fetch_command_usage_text      string = "to fetch templates use fetch command with desired template path separated by /"
	fetch_command_usage_text_repo string = fmt.Sprintf("to fetch [DOC] template use: %s %s --%s [path/to/repo] --%s [folder name](optional)", command_name, command_name_fetch, arg_name_repo, arg_name_name)
	fetch_command_usage_text_doc  string = fmt.Sprintf("to fetch [DOC] template use: %s %s --%s [path/to/file] --%s [file name](optional)", command_name, command_name_fetch, arg_name_doc, arg_name_name)
)

func printFetchCommandFullExample() {
	fmt.Println(fetch_command_usage_text)
	fmt.Println(fetch_command_usage_text_repo)
	fmt.Println(fetch_command_usage_text_doc)
}
