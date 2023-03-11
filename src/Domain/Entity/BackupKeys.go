package entity

import (
	"encoding/json"

	iamValueObject "../ValueObject/Iam"
)

type BackupKeys struct {
	accessKeyId     *iamValueObject.AccessKeyId
	secretAccessKey *iamValueObject.SecretAccessKey
}

func NewBackupKeys(
	accessKeyId *iamValueObject.AccessKeyId,
	secretAccessKey *iamValueObject.SecretAccessKey,
) *BackupKeys {
	return &BackupKeys{
		accessKeyId:     accessKeyId,
		secretAccessKey: secretAccessKey,
	}
}

func (self BackupKeys) GetAccessKeyId() *iamValueObject.AccessKeyId {
	return self.accessKeyId
}

func (self BackupKeys) GetSecretAccessKey() *iamValueObject.SecretAccessKey {
	return self.secretAccessKey
}

func (self BackupKeys) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]string{
		"accessKeyId":     self.accessKeyId.String(),
		"secretAccessKey": self.secretAccessKey.String(),
	})
}
