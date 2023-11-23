package configuration

type Config struct {
	User     string
	Password string
	Host     string
	Port     string
	Database string
}

func AzureAccountDBConfig() *Config {
	return &Config{
		Host:     "jeab-test-server.postgres.database.azure.com",
		Database: "account_db",
		User:     "amVhYi10ZXN0LXNlcnZlcg",
		Password: "amd1YXJk",
	}
}
