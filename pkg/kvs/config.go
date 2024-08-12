package kvs

import "golang-example/pkg/credentials"

type Config struct {
	TableName   string `env:"PAYOUT_TABLE"`
	AWSEndpoint string `env:"AWS_ENDPOINT"`
	Credential  credentials.Credential
}
