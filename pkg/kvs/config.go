package kvs

import "golang-example/pkg/credential"

type Config struct {
	TableName  string `env:"TABLE_NAME_PROJECTION"`
	Credential credential.Credential
}
