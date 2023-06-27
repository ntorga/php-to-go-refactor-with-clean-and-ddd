package iamValueObject

import (
	"errors"
	"regexp"
)

type AccessKeyId string

func NewAccessKeyId(value string) (AccessKeyId, error) {
	aki := AccessKeyId(value)
	if !aki.isValid() {
		return "", errors.New("InvalidAccessKeyId")
	}
	return aki, nil
}

func NewAccessKeyIdPanic(value string) AccessKeyId {
	aki := AccessKeyId(value)
	if !aki.isValid() {
		panic("InvalidAccessKeyId")
	}
	return aki
}

func (aki AccessKeyId) isValid() bool {
	re := regexp.MustCompile(`^(?<![A-Z0-9])[A-Z0-9]{20}(?![A-Z0-9])$`)
	return re.MatchString(string(aki))
}

func (aki AccessKeyId) String() string {
	return string(aki)
}
