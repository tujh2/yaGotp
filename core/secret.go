package core

import "encoding/base32"

type Secret struct {
	data []byte
}

func (secret *Secret) CreateSecret(secretString string) bool {
	data, err := base32.StdEncoding.WithPadding(base32.NoPadding).DecodeString(secretString)
	if err != nil {
		println(err.Error())
		return false
	}
	secret.data = data[0 : 16]
	return true
}

func (secret *Secret) GetBytes() []byte {
	return secret.data
}
