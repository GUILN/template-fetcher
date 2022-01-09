package configuration

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/guiln/boilerplate-cli/helpers"
)

const (
	configFileName    string = "fetcher.config"
	templatesFileName string = "templates.json"
)

type Config struct {
	Repo              string `json:"repo"`
	RepoOwner         string `json:"repo_owner"`
	Token             string `json:"token"`
	secret            string
	configDirPath     string
	configFilePath    string
	templatesFilePath string
}

func NewConfig(configDirPath, secret string) *Config {
	return &Config{
		secret:            secret,
		configDirPath:     configDirPath,
		configFilePath:    filepath.Join(configDirPath, configFileName),
		templatesFilePath: filepath.Join(configDirPath, templatesFileName),
	}
}

// SaveToken saves encrypted token in the config
func (c *Config) GetToken() (string, *ConfigError) {
	decryptedToken, err := helpers.Decrypt(c.Token, c.secret)
	if err != nil {
		return "", CreateConfigError("error when trying to decrypt token", err)
	}

	return decryptedToken, nil
}

func (c *Config) SetToken(token string) *ConfigError {
	encryptedToken, err := helpers.Encrypt(token, c.secret)
	if err != nil {
		return CreateConfigError("enrror when trying to encrypt git token", err)
	}

	c.Token = encryptedToken
	return nil
}

func (c *Config) LoadConfig() *ConfigError {
	if !helpers.CheckPathExists(c.configDirPath) {
		if err := os.Mkdir(c.configDirPath, 0755); err != nil {
			return CreateConfigError(fmt.Sprintf("error while trying create %s dir", c.configDirPath), err)
		}
	}
	cFile, err := ioutil.ReadFile(c.configFilePath)
	if err != nil {
		return CreateConfigError(fmt.Sprintf("error while trying to read config file at %s", c.configDirPath), err)
	}
	if err := json.Unmarshal(cFile, c); err != nil {
		return CreateConfigError("error while trying to unmarshal cofig file", err)
	}

	return nil
}

func (c *Config) PersistConfig() *ConfigError {

	jsonString, err := json.Marshal(c)
	if err != nil {
		return CreateConfigError("error while trying to unmarshal config to be persisted", err)
	}
	err = ioutil.WriteFile(c.configFilePath, []byte(jsonString), os.ModePerm)
	if err != nil {
		return CreateConfigError(fmt.Sprintf("error while trying to write config to file %s", c.configDirPath), err)
	}
	return nil
}

func (c *Config) String() string {
	gitToken, _ := c.GetToken()
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("Repo:           %s\n", c.Repo))
	builder.WriteString(fmt.Sprintf("RepoOwner:      %s\n", c.RepoOwner))
	builder.WriteString(fmt.Sprintf("GitToken:       %s\n", gitToken))

	return builder.String()
}

type ConfigError struct {
	message    string
	innerError error
}

func (ce *ConfigError) Error() string {
	builder := strings.Builder{}
	builder.WriteString("ConfigErrorMessage:")
	builder.WriteString(ce.message)
	builder.WriteString("InnerErrorMessage:")
	builder.WriteString(ce.innerError.Error())
	return builder.String()
}

func CreateConfigError(message string, err error) *ConfigError {
	return &ConfigError{message: message, innerError: err}
}
