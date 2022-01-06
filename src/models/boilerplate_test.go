package models

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GivenIHaveAFolderStructureInBoilerplateFolder_WhenICallMarshalFunction_ThenItReturnsExpectedString(t *testing.T) {
	someBoilerplateRepo := getBoilerplateRepo()

	marshalledRoot, err := someBoilerplateRepo.RootBoilerplateFolder.Marshal()

	assert.Nil(t, err)
	assert.Contains(t, marshalledRoot, "node/rest-api")

	fmt.Println(marshalledRoot)
}

func getBoilerplateRepo() *BoilerplateRepo {

	rootFolder := NewBoilerplateFolder("", false)
	rootFolder.AddChild(&BoilerplateFolder{
		isContainerFolder: false,
		isRootFolder:      false,
		path:              "node",
		childBoilerplateFolders: []*BoilerplateFolder{
			&BoilerplateFolder{
				isContainerFolder:       true,
				isRootFolder:            false,
				path:                    "node/rest-api",
				childBoilerplateFolders: []*BoilerplateFolder{},
			},
		},
	}, &BoilerplateFolder{
		isContainerFolder: false,
		isRootFolder:      false,
		path:              "go",
		childBoilerplateFolders: []*BoilerplateFolder{
			&BoilerplateFolder{
				isContainerFolder:       true,
				isRootFolder:            false,
				path:                    "go/rest-api",
				childBoilerplateFolders: []*BoilerplateFolder{},
			},
			&BoilerplateFolder{
				isContainerFolder:       true,
				isRootFolder:            false,
				path:                    "go/cli",
				childBoilerplateFolders: []*BoilerplateFolder{},
			},
		},
	})
	boilerplateRepo := NewBoilerplateRepo(rootFolder)
	return boilerplateRepo
}
