package configuration

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/guiln/boilerplate-cli/src/crypto"
)

const (
	configDirPath     string = "~/.template.fetcher"
	configFileName    string = "fetcher.config"
	templatesFileName string = "templates.json"
)

var (
	configFilePath    string = filepath.Join(configDirPath, configFileName)
	templatesFilePath string = filepath.Join(configDirPath, templatesFileName)
)

type Config struct {
	Repo      string `json:"repo"`
	RepoOwner string `json:"repo_owner"`
	token     string `json:"token"`
	secret    string
}

func NewConfig(repoName, repoOwner, secret string) *Config {
	return &Config{
		Repo:      repoName,
		RepoOwner: repoOwner,
		secret:    secret,
	}
}

// SaveToken saves encrypted token in the config
func (c *Config) GetToken() (string, *ConfigError) {
	decryptedToken, err := crypto.Decrypt(c.token, c.secret)
	if err != nil {
		return "", CreateConfigError("error when trying to decrypt token", err)
	}

	return decryptedToken, nil
}

func (c *Config) SetToken(token string) *ConfigError {
	encryptedToken, err := crypto.Encrypt(token, c.secret)
	if err != nil {
		return CreateConfigError("enrror when trying to encrypt git token", err)
	}

	c.token = encryptedToken
	return nil
}

func (c *Config) LoadConfig() *ConfigError {
	if !checkPathExists(configDirPath) {
		if err := os.Mkdir(configDirPath, 0755); err != nil {
			return CreateConfigError(fmt.Sprintf("error ocurred when trying to create folder %s", configDirPath), err)
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
