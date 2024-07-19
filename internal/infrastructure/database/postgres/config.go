package postgres

type PostgresConnectionConfig struct {
	Host     string
	User     string
	Password string
	DbName   string
	Port     string
	SslMode  string
	TimeZone string
}
