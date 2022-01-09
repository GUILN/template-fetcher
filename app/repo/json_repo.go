package repo

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/guiln/boilerplate-cli/src/models"
)

type JsonTemplateRepo struct {
	jsonPath string
}

func NewTemplateJsonRepo(jsonPath string) *JsonTemplateRepo {
	return &JsonTemplateRepo{jsonPath: jsonPath}
}

func (jtr *JsonTemplateRepo) ReadRepo() (*models.BoilerplateRepo, *models.BoilerplateError) {
	file, err := ioutil.ReadFile(jtr.jsonPath)
	if err != nil {
		return nil, models.CreateBoilerplateErrorFromError(err, fmt.Sprintf("error while trying to create file on path %s", jtr.jsonPath))
	}
	rootFolder, bErr := models.JsonUnmarshal(string(file))
	if err != nil {
		return nil, bErr
	}

	bRepo := models.NewBoilerplateRepo(rootFolder)

	return bRepo, nil
}

func (jtr *JsonTemplateRepo) PersistRepo(bRepo *models.BoilerplateRepo) *models.BoilerplateError {
	rootFolder, err := bRepo.RootBoilerplateFolder.JsonMarshal()
	if err != nil {
		return err
	}
	if err := ioutil.WriteFile(jtr.jsonPath, []byte(rootFolder), os.ModePerm); err != nil {
		return models.CreateBoilerplateErrorFromError(err, fmt.Sprintf("error while trying to write template file on path %s", jtr.jsonPath))
	}
	return nil
}
