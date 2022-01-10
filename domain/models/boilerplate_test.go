package models

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

const boilerPlateRepoJson string = `{"IsContainerFolder":false,"is_root_folder":true,"path":"TEMPLATES","child_boilerplate_folders":[{"IsContainerFolder":false,"is_root_folder":false,"path":"node","child_boilerplate_folders":[{"IsContainerFolder":true,"is_root_folder":false,"path":"node/rest-api","child_boilerplate_folders":[]}]},{"IsContainerFolder":false,"is_root_folder":false,"path":"go","child_boilerplate_folders":[{"IsContainerFolder":true,"is_root_folder":false,"path":"go/rest-api","child_boilerplate_folders":[]},{"IsContainerFolder":true,"is_root_folder":false,"path":"go/cli","child_boilerplate_folders":[]}]}]}`

func Test_GivenIHaveAFolderStructureInBoilerplateFolder_WhenICallStringFunction_ThenItReturnsExpectedString(t *testing.T) {
	someBoilerplateRepo := getBoilerplateRepo()

	marshalledRoot, err := someBoilerplateRepo.RootBoilerplateFolder.String()

	expected := "TEMPLATES\n   node\n      rest-api\n   go\n      rest-api\n      cli"
	assert.Nil(t, err)
	assert.Equal(t, expected, marshalledRoot)

	fmt.Println(marshalledRoot)
}

func Test_GivenAFolderStructure_WhenICallJsonMarshalFunction_ThenItReturnsExpectedJsonStructure(t *testing.T) {
	someRepo := getBoilerplateRepo()

	jsonRepo, err := someRepo.RootBoilerplateFolder.JsonMarshal()

	assert.Nil(t, err)
	fmt.Println(jsonRepo)
}

func Test_GivenAJsonRepresentingFolderStructure_WhenICallJsonUnmarshalFunction_ThenItReturnsExpectedBoilerplateFolderStructure(t *testing.T) {
	rootFolder, err := JsonUnmarshal(boilerPlateRepoJson)

	assert.Nil(t, err)
	assert.NotNil(t, rootFolder)
	assert.NotNil(t, rootFolder.ChildBoilerplateFolders)
	assert.Equal(t, "TEMPLATES", rootFolder.Path)
}

func getBoilerplateRepo() *BoilerplateRepo {

	rootFolder := NewBoilerplateFolder("TEMPLATES", false)
	rootFolder.AddChild(&BoilerplateFolder{
		IsContainerFolder: false,
		IsRootFolder:      false,
		Path:              "node",
		ChildBoilerplateFolders: []*BoilerplateFolder{
			&BoilerplateFolder{
				IsContainerFolder:       true,
				IsRootFolder:            false,
				Path:                    "node/rest-api",
				ChildBoilerplateFolders: []*BoilerplateFolder{},
			},
		},
	}, &BoilerplateFolder{
		IsContainerFolder: false,
		IsRootFolder:      false,
		Path:              "go",
		ChildBoilerplateFolders: []*BoilerplateFolder{
			&BoilerplateFolder{
				IsContainerFolder:       true,
				IsRootFolder:            false,
				Path:                    "go/rest-api",
				ChildBoilerplateFolders: []*BoilerplateFolder{},
			},
			&BoilerplateFolder{
				IsContainerFolder:       true,
				IsRootFolder:            false,
				Path:                    "go/cli",
				ChildBoilerplateFolders: []*BoilerplateFolder{},
			},
		},
	})
	boilerplateRepo := NewBoilerplateRepo(rootFolder)
	return boilerplateRepo
}
