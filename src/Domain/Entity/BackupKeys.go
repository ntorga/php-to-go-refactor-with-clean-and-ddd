package entity

import (
	iamValueObject "../ValueObject/Iam"
)

type BackupKeys struct {
	AccessKeyId     iamValueObject.AccessKeyId     `json:"accessKeyId"`
	SecretAccessKey iamValueObject.SecretAccessKey `json:"secretAccessKey"`
}

func NewBackupKeys(
	accessKeyId iamValueObject.AccessKeyId,
	secretAccessKey iamValueObject.SecretAccessKey,
) BackupKeys {
	return BackupKeys{
		AccessKeyId:     accessKeyId,
		SecretAccessKey: secretAccessKey,
	}
}

func (bk BackupKeys) GetAccessKeyId() iamValueObject.AccessKeyId {
	return bk.AccessKeyId
}

func (bk BackupKeys) GetSecretAccessKey() iamValueObject.SecretAccessKey {
	return bk.SecretAccessKey
}
