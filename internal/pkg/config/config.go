package config

import "github.com/kelseyhightower/envconfig"

// DBConfiguration - config used for DynamoDB
type DBConfiguration struct {
	AccessKey   string `envconfig:"DB_AWS_ACCESS_KEY_ID" required:"true"`
	SecretKey   string `envconfig:"DB_AWS_SECRET_ACCESS_KEY" required:"true"`
	Region      string `envconfig:"DB_AWS_DEFAULT_REGION" required:"true"`
	EndpointURL string `envconfig:"DB_AWS_ENDPOINT_URL" required:"true"`
	TableName   string `envconfig:"DB_TABLE_NAME" required:"true"`
}

// GetDBConfig get db env vars or error
func GetDBConfig() (DBConfiguration, error) {
	dbConfig := DBConfiguration{}
	err := envconfig.Process("", &dbConfig)
	if err != nil {
		return dbConfig, err
	}
	return dbConfig, nil
}
