package credentials

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
)

type Credential interface {
	GetConfig() aws.Config
}

type credentialAWS struct {
	awsConfig aws.Config
}

func NewCredential(ctx context.Context, cfg ConfigCredential) (Credential, error) {
	awsCredential, err := config.LoadDefaultConfig(ctx,
		config.WithRegion(cfg.AWSRegion),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			cfg.AWSAccessKey,
			cfg.AWSSecretKey,
			"")),
	)

	/*awsCredential, err := New(ctx,
	WithEndpoint(cfg.AWSEndpoint),
	WithRegion(cfg.AWSRegion),
	WithRole(cfg.RoleARN))*/
	if err != nil {
		return &credentialAWS{}, err
	}
	return &credentialAWS{
		awsConfig: awsCredential,
	}, err
}

func (c *credentialAWS) GetConfig() aws.Config {
	return c.awsConfig
}
