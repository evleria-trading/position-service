package config

type Сonfig struct {
	Environment string `env:"ENVIRONMENT" envDefault:"dev"`

	PostgresUser       string `env:"POSTGRES_USER" envDefault:"postgres"`
	PostgresPass       string `env:"POSTGRES_PASSWORD" envDefault:""`
	PostgresHost       string `env:"POSTGRES_HOST" envDefault:"localhost"`
	PostgresPort       int    `env:"POSTGRES_PORT" envDefault:"5432"`
	PostgresDb         string `env:"POSTGRES_DB" envDefault:"positions_db"`
	PostgresSSLDisable bool   `env:"POSTGRES_SSL_DISABLE" envDefault:"false"`

	PriceServiceHost string `env:"PRICE_SERVICE_HOST" envDefault:"localhost"`
	PriceServicePort int    `env:"PRICE_SERVICE_PORT" envDefault:"6000"`

	UserServiceHost string `env:"USER_SERVICE_HOST" envDefault:"localhost"`
	UserServicePort int    `env:"USER_SERVICE_PORT" envDefault:"6000"`
}
