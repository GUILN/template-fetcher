package crypto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	mySuperSecretString  string = "My_Super_SecretString"
	expectedEncryptedStr string = "Jjl40B1cRuo+UcrwYSyQMggbtXUK"
	mySecret             string = "abc&1*~#^2^#s0^=)^^7%b34"
)

func Test_GivenICallEncryptFunction_WhenIProvideAString_ThenIGetEncryptedString(t *testing.T) {
	encStr, err := Encrypt(mySuperSecretString, mySecret)

	assert.Nil(t, err)
	assert.Equal(t, expectedEncryptedStr, encStr)
}

func Test_GivenICallDecryptFunction_WhenIProvideEncryptedString_ThenIGetOriginalString(t *testing.T) {
	decodedStr, err := Decrypt(expectedEncryptedStr, mySecret)

	assert.Nil(t, err)
	assert.Equal(t, mySuperSecretString, decodedStr)
}
