package repo

import "github.com/guiln/boilerplate-cli/src/models"

type TemplateJsonRepo struct {
	jsonPath string
}

func NewTemplateJsonRepo(jsonPath string) *TemplateJsonRepo {
	return &TemplateJsonRepo{jsonPath: jsonPath}
}

func (tjr *TemplateJsonRepo) ReadRepo() (*models.BoilerplateRepo, error) {
	// TODO
	return nil, nil
}

func (tjr *TemplateJsonRepo) SaveRepo(bRepo *models.BoilerplateRepo) error {
	// TODO
	return nil
}
