package models

import (
	"encoding/json"
	"path/filepath"
	"strings"
)

// BoilerplateRepo holds the pointer to templates root folder
type BoilerplateRepo struct {
	RootBoilerplateFolder *BoilerplateFolder `json:"RootTemplate"`
}

// NewBoilerplateRepo has the same effect as creating directly through &BoilerplateRepo{}
// But setting isContainerFolder = false && isRootFolder = true by default
func NewBoilerplateRepo(rootFolder *BoilerplateFolder) *BoilerplateRepo {
	rootFolder.IsContainerFolder = false
	rootFolder.IsRootFolder = true

	return &BoilerplateRepo{RootBoilerplateFolder: rootFolder}
}

// Holds the structure and info of a template folder with nested
type BoilerplateFolder struct {
	IsContainerFolder       bool                 `json:is_container_folder"`
	IsRootFolder            bool                 `json:"is_root_folder"`
	Path                    string               `json:"path"`
	ChildBoilerplateFolders []*BoilerplateFolder `json:"child_boilerplate_folders"`
}

func NewBoilerplateFolder(folderPath string, isContainer bool) *BoilerplateFolder {
	return &BoilerplateFolder{Path: folderPath, IsContainerFolder: isContainer, IsRootFolder: false}
}

func JsonUnmarshal(jsonString string) (*BoilerplateFolder, *BoilerplateError) {
	var rootFolder BoilerplateFolder
	err := json.Unmarshal([]byte(jsonString), &rootFolder)
	if err != nil {
		return nil, CreateBoilerplateErrorFromError(err, "error while trying to unmarshal boilerplate folder json")
	}
	return &rootFolder, nil
}

func (bfolder *BoilerplateFolder) String() (string, *BoilerplateError) {
	return printTree(bfolder, 0), nil
}

func (bfolder *BoilerplateFolder) JsonMarshal() (string, *BoilerplateError) {
	outXml, err := json.Marshal(bfolder)
	if err != nil {
		return "", CreateBoilerplateErrorFromError(err, "error when trying to marshal boilerplate folder to xml")
	}

	return string(outXml), nil
}

func (bFolder *BoilerplateFolder) GetPath() string {
	return bFolder.Path
}

func (bFolder *BoilerplateFolder) SetPath(path string) {
	bFolder.Path = path
}

func (bFolder *BoilerplateFolder) AddChild(boilerplateFolders ...*BoilerplateFolder) {
	bFolder.ChildBoilerplateFolders = append(bFolder.ChildBoilerplateFolders, boilerplateFolders...)
}

func (bFolder *BoilerplateFolder) SetIsContainer(isContainer bool) {
	bFolder.IsContainerFolder = isContainer
}

func (bFolder *BoilerplateFolder) IsContainer() bool {
	return bFolder.IsContainerFolder
}

func (bFolder *BoilerplateFolder) IsRoot() bool {
	return bFolder.IsRootFolder
}

func printTree(folder *BoilerplateFolder, level int) string {
	baseFolder := filepath.Base(folder.GetPath())
	var tree string
	if baseFolder != "/" {
		level++
		tree = strings.Repeat(" ", level*3) + baseFolder
	} else {
		if level == 0 {
			tree = "TEMPLATES:"
		}
	}
	for _, child := range folder.ChildBoilerplateFolders {
		tree += "\n"
		tree += printTree(child, level)
	}

	return tree
}
