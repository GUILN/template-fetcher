package models

import (
	"encoding/json"
	"fmt"
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
	rootFolder.IsRepoContainerFolder = false
	rootFolder.IsRootFolder = true

	return &BoilerplateRepo{RootBoilerplateFolder: rootFolder}
}

// Holds the structure and info of a template folder with nested
type BoilerplateFolder struct {
	IsRepoContainerFolder   bool                 `json:is_repo_container_folder"`
	IsDocContainerFolder    bool                 `json:is_doc_container_folder"`
	IsRootFolder            bool                 `json:"is_root_folder"`
	Path                    string               `json:"path"`
	ChildBoilerplateFolders []*BoilerplateFolder `json:"child_boilerplate_folders"`
	ChildTemplateDocuments  []*TemplateDocument  `json:child_template_documents`
}

func NewBoilerplateFolder(folderPath string, isRepoContainer bool, isDocContainer bool) *BoilerplateFolder {
	return &BoilerplateFolder{Path: folderPath, IsRepoContainerFolder: isRepoContainer, IsDocContainerFolder: isDocContainer, IsRootFolder: false}
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

func (bFolder *BoilerplateFolder) AddChildDoc(templateDoc ...*TemplateDocument) *BoilerplateError {
	if !bFolder.IsDocContainerFolder {
		return &BoilerplateError{Message: "cannot add template docs into non doc container folder"}
	}
	bFolder.ChildTemplateDocuments = append(bFolder.ChildTemplateDocuments, templateDoc...)
	return nil
}

func (bFolder *BoilerplateFolder) SetIsRepoContainer(isContainer bool) {
	bFolder.IsRepoContainerFolder = isContainer
}

func (bFolder *BoilerplateFolder) IsRepoContainer() bool {
	return bFolder.IsRepoContainerFolder
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
		if folder.IsRepoContainerFolder {
			tree += " [REPO]"
		} else if folder.IsDocContainerFolder {
			tree += "\n" + printChildDocs(folder, strings.Repeat(" ", (level+1)*3))
		}
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

func printChildDocs(folder *BoilerplateFolder, space string) string {
	if len(folder.ChildTemplateDocuments) == 0 {
		panic(fmt.Errorf("cannot print child template documents if there is no template documents"))
	}
	sBuilder := strings.Builder{}
	for _, tdoc := range folder.ChildTemplateDocuments {
		sBuilder.WriteString(space)
		sBuilder.WriteString(tdoc.Name)
		sBuilder.WriteString(" [DOC]")
		sBuilder.WriteString("\n")
	}
	return sBuilder.String()
}

type TemplateDocument struct {
	Name string `json:"name"`
	Path string `json:"path"`
}
