package configuration

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	mySecret       string = "abc&1*~#^2^#s0^=)^^7%b34"
	gitToken       string = "my_git_token"
	encryptedToken string = "Bjl45AFYfOwOacr9"
)

func Test_GivenISetAndGetToken_ThenTokenIsSetEncryptedAndGotDecrypted(t *testing.T) {
	cfg := NewConfig("repoName", "repoOwner", mySecret)
	err := cfg.SetToken(gitToken)

	assert.Nil(t, err)
	assert.Equal(t, encryptedToken, cfg.token)

	tkn, err := cfg.GetToken()
	assert.Nil(t, err)
	assert.Equal(t, gitToken, tkn)
}
