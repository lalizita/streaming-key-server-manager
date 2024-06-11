package env

type EnvConfig struct {
	PostgresHost string `env:"POSTGRES_HOST"`
	PostgresUser string `env:"POSTGRES_USER"`
	PostgrePass  string `env:"POSTGRES_PASSWORD"`
	PostgresDB   string `env:"POSTGRES_DB"`
	PostgresPort string `env:"POSTGRES_PORT"`
}
