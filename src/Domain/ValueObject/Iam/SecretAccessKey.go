package iamValueObject

import (
	"encoding/json"
	"regexp"
)

type SecretAccessKey struct {
	secretAccessKey string
}

func NewSecretAccessKey(secretAccessKey string) *SecretAccessKey {
	if !isValidSecretAccessKey(secretAccessKey) {
		panic("InvalidSecretAccessKey")
	}
	return &SecretAccessKey{secretAccessKey: secretAccessKey}
}

func (self *SecretAccessKey) String() string {
	return self.secretAccessKey
}

func (self *SecretAccessKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(self.secretAccessKey)
}

func isValidSecretAccessKey(secretAccessKey string) bool {
	pattern := `^(?<![A-Za-z0-9\/+=])[A-Za-z0-9\/+=]{40}(?![A-Za-z0-9\/+=])$`
	matched, _ := regexp.MatchString(pattern, secretAccessKey)
	return matched
}
