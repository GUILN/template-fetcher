package models

import "strings"

// BoilerplateRepo holds the pointer to templates root folder
type BoilerplateRepo struct {
	RootBoilerplateFolder *BoilerplateFolder `yaml:"templates"`
}

// NewBoilerplateRepo has the same effect as creating directly through &BoilerplateRepo{}
// But setting isContainerFolder = false && isRootFolder = true by default
func NewBoilerplateRepo(rootFolder *BoilerplateFolder) *BoilerplateRepo {
	rootFolder.isContainerFolder = false
	rootFolder.isRootFolder = true

	return &BoilerplateRepo{RootBoilerplateFolder: rootFolder}
}

// Holds the structure and info of a template folder with nested
type BoilerplateFolder struct {
	isContainerFolder       bool
	isRootFolder            bool
	path                    string
	childBoilerplateFolders []*BoilerplateFolder
}

func NewBoilerplateFolder(folderPath string, isContainer bool) *BoilerplateFolder {
	return &BoilerplateFolder{path: folderPath, isContainerFolder: isContainer, isRootFolder: false}
}

func (bfolder *BoilerplateFolder) Marshal() (string, *BoilerplateError) {
	return printTree(bfolder, 0), nil
}

func (bFolder *BoilerplateFolder) GetPath() string {
	return bFolder.path
}

func (bFolder *BoilerplateFolder) SetPath(path string) {
	bFolder.path = path
}

func (bFolder *BoilerplateFolder) AddChild(boilerplateFolders ...*BoilerplateFolder) {
	bFolder.childBoilerplateFolders = append(bFolder.childBoilerplateFolders, boilerplateFolders...)
}

func (bFolder *BoilerplateFolder) SetIsContainer(isContainer bool) {
	bFolder.isContainerFolder = isContainer
}

func (bFolder *BoilerplateFolder) IsContainer() bool {
	return bFolder.isContainerFolder
}

func (bFolder *BoilerplateFolder) IsRoot() bool {
	return bFolder.isRootFolder
}

func printTree(folder *BoilerplateFolder, level int) string {
	tree := strings.Repeat(" ", level) + "|" + "\n"
	tree += strings.Repeat("-", level) + folder.GetPath()

	for _, child := range folder.childBoilerplateFolders {
		tree += "\n"
		tree += printTree(child, level+1)
	}

	return tree
}
