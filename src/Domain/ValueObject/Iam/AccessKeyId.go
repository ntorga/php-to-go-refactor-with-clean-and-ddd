package iamValueObject

import (
	"encoding/json"
	"regexp"
)

type AccessKeyId struct {
	accessKeyId string
}

func NewAccessKeyId(accessKeyId string) *AccessKeyId {
	if !isValidAccessKeyId(accessKeyId) {
		panic("InvalidAccessKeyId")
	}
	return &AccessKeyId{accessKeyId: accessKeyId}
}

func (self *AccessKeyId) String() string {
	return self.accessKeyId
}

func (self *AccessKeyId) MarshalJSON() ([]byte, error) {
	return json.Marshal(self.accessKeyId)
}

func isValidAccessKeyId(accessKeyId string) bool {
	pattern := `^(?<![A-Z0-9])[A-Z0-9]{20}(?![A-Z0-9])$`
	matched, _ := regexp.MatchString(pattern, accessKeyId)
	return matched
}
