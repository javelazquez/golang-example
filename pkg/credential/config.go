package credential

type ConfigCredential struct {
	AWSRegion    string `env:"AWS_REGION"`
	AWSAccessKey string `env:"AWS_ACCESS_KEY_ID"`
	AWSSecretKey string `env:"AWS_SECRET_ACCESS_KEY"`
}
