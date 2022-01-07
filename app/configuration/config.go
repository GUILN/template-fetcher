package configuration

import (
	"fmt"
	"os"
	"path/filepath"
)

const (
	templateFetcherPath        string = "~/.template.fetcher"
	configFileName             string = "fetcher.config"
	availableTemplatesFileName string = "templates"
)

var (
	configFilePath             string = filepath.Join(templateFetcherPath, configFileName)
	availableTemplatesFilePath string = filepath.Join(templateFetcherPath, availableTemplatesFileName)
)

type Config struct {
	Repo      string `json:"repo"`
	token     string `json:"token"`
	RepoOwner string `json:"repo_owner"`
}

func (c *Config) NewConfig(repoName, repoOwner string) *Config {
	return &Config{
		Repo:      repoName,
		RepoOwner: repoOwner,
	}
}

// SaveToken saves encrypted token in the config
func (c *Config) GetToken() string {
	//TODO: decrypt the token
	return "" // return token
}

func (c *Config) SetToken(token string) *ConfigError {
	// TODO encrypt the token
	// TODO sets toke like: c.token = encryptedToken
	return nil
}

func (c *Config) LoadConfig() *ConfigError {
	if !checkPathExists(templateFetcherPath) {
		if err := os.Mkdir(templateFetcherPath, 0755); err != nil {
			return CreateConfigError(fmt.Sprintf("error ocurred when trying to create folder %s", templateFetcherPath), err)
		}
	}
	if !checkPathExists(configFilePath) {
		return CreateConfigError("could not find config file, make sure config was created by running any config command", nil)
	}
	// TODO: load config from file here
	// TODO: decrypt token
	return nil
}

type ConfigError struct {
	message    string
	innerError error
}

func (ce *ConfigError) Error() string {
	return ce.message
}

func CreateConfigError(message string, err error) *ConfigError {
	return &ConfigError{message: message, innerError: err}
}

func checkPathExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}

	return true
}
