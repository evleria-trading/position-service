package config

type Сonfig struct {
	RedisPass string `env:"REDIS_PASS" envDefault:""`
	RedisHost string `env:"REDIS_HOST" envDefault:"localhost"`
	RedisPort int    `env:"REDIS_PORT" envDefault:"6379"`

	PostgresUser       string `env:"POSTGRES_USER" envDefault:"postgres"`
	PostgresPass       string `env:"POSTGRES_PASSWORD" envDefault:""`
	PostgresHost       string `env:"POSTGRES_HOST" envDefault:"127.0.0.1"`
	PostgresPort       int    `env:"POSTGRES_PORT" envDefault:"5432"`
	PostgresDb         string `env:"POSTGRES_DB" envDefault:"postgres"`
	PostgresSSLDisable bool   `env:"POSTGRES_SSL_DISABLE" envDefault:"false"`
}
