package models

type BoilerplateRepo struct {
	RootBoilerplateFolder *BoilerplateFolder
}

type BoilerplateFolder struct {
	isContainerFolder       bool
	childBoilerplateFolders []*BoilerplateFolder
}
