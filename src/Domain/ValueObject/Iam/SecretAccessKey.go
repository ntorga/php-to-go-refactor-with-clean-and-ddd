package iamValueObject

import (
	"errors"
	"regexp"
)

type SecretAccessKey string

func NewSecretAccessKey(value string) (SecretAccessKey, error) {
	sak := SecretAccessKey(value)
	if !sak.isValid() {
		return "", errors.New("InvalidSecretAccessKey")
	}
	return sak, nil
}

func NewSecretAccessKeyPanic(value string) SecretAccessKey {
	sak := SecretAccessKey(value)
	if !sak.isValid() {
		panic("InvalidSecretAccessKey")
	}
	return sak
}

func (sak SecretAccessKey) isValid() bool {
	re := regexp.MustCompile(`^(?<![A-Za-z0-9\/+=])[A-Za-z0-9\/+=]{40}(?![A-Za-z0-9\/+=])$`)
	return re.MatchString(string(sak))
}

func (sak SecretAccessKey) String() string {
	return string(sak)
}
