package cryptof

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAes_CFB(t *testing.T) {
	var (
		aes  = NewAes()
		key  = []byte("light88888888888")
		data = "什么鬼！？"
	)

	ciphertext, err := aes.CFBEncrypt(data, key)
	if err != nil {
		t.Error(err)
	}

	plaintext, err := aes.CFBDecrypt(ciphertext, key)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, plaintext, data)
}

func TestAes_CBC(t *testing.T) {
	var (
		aes  = NewAes()
		key  = []byte("light88888888888")
		data = "什么鬼！？"
	)

	ciphertext, err := aes.CBCEncrypt(data, key, nil)
	if err != nil {
		t.Error(err)
	}

	plaintext, err := aes.CBCDecrypt(ciphertext, key, nil)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, plaintext, data)
}
