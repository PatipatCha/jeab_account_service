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
		Host:     "jeab-stg-server.postgres.database.azure.com",
		Database: "operation_center_db",
		User:     "amVhYi10ZXN0LXNlcnZlc",
		Password: "MTIvMDEvMjAyM0FBQQ",
	}
}
