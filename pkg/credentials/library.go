package credentials

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials/stscreds"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

type awsConfigOpts struct {
	region   string
	role     string
	endpoint string
	profile  string
}

// New builds an aws config with the options received
func New(ctx context.Context, opts ...Option) (aws.Config, error) {

	configOpts := awsConfigOpts{}

	for _, o := range opts {
		o(&configOpts)
	}

	return build(ctx, configOpts)
}

// Option is a function to add a new custom property to the aws config
type Option func(configOpts *awsConfigOpts)

// WithRegion option to add a region for aws config
func WithRegion(region string) Option {
	return func(configOpts *awsConfigOpts) {
		configOpts.region = region
	}
}

// WithRole option to add a role for aws config
func WithRole(role string) Option {
	return func(configOpts *awsConfigOpts) {
		configOpts.role = role
	}
}

// WithEndpoint option to add an endpoint for aws config
func WithEndpoint(endpoint string) Option {
	return func(configOpts *awsConfigOpts) {
		configOpts.endpoint = endpoint
	}
}

// WithProfile option to add profile credentials
func WithProfile(profile string) Option {
	return func(configOpts *awsConfigOpts) {
		configOpts.profile = profile
	}
}

func build(ctx context.Context, opts awsConfigOpts) (aws.Config, error) {
	if opts.region == "" {
		opts.region = "eu-west-1"
	}

	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(opts.region), config.WithSharedConfigProfile(opts.profile))
	if err != nil {
		return aws.Config{}, err
	}
	if opts.role != "" {
		cfg = SetupRole(cfg, opts.role)
	}
	if opts.endpoint != "" {
		cfg = SetupEndpoint(cfg, opts.endpoint)
	}

	return cfg, nil
}

// EmptyConfig loads a new empty aws configs
func EmptyConfig(ctx context.Context, region string) (aws.Config, error) {
	return config.LoadDefaultConfig(ctx, config.WithRegion(region))
}

// SetupRole configures a role of an aws config
func SetupRole(cfg aws.Config, role string) aws.Config {
	cfg.Credentials = aws.NewCredentialsCache(stscreds.NewAssumeRoleProvider(
		sts.NewFromConfig(cfg),
		role,
	))
	return cfg
}

// SetupEndpoint setups the aws conf to use a custom endpoint
func SetupEndpoint(cfg aws.Config, endpoint string) aws.Config {
	cfg.EndpointResolverWithOptions = aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{
			PartitionID:   "aws",
			URL:           endpoint,
			SigningRegion: region,
		}, nil
	})
	return cfg
}
