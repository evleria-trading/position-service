package config

import "time"

type Сonfig struct {
	Environment string `env:"ENVIRONMENT" envDefault:"dev"`

	RedisPass string `env:"REDIS_PASS" envDefault:""`
	RedisHost string `env:"REDIS_HOST" envDefault:"localhost"`
	RedisPort int    `env:"REDIS_PORT" envDefault:"6379"`

	PostgresUser       string `env:"POSTGRES_USER" envDefault:"postgres"`
	PostgresPass       string `env:"POSTGRES_PASSWORD" envDefault:""`
	PostgresHost       string `env:"POSTGRES_HOST" envDefault:"localhost"`
	PostgresPort       int    `env:"POSTGRES_PORT" envDefault:"5432"`
	PostgresDb         string `env:"POSTGRES_DB" envDefault:"postgres"`
	PostgresSSLDisable bool   `env:"POSTGRES_SSL_DISABLE" envDefault:"false"`

	GeneratePrices bool          `env:"GENERATE_PRICES" envDefault:"true"`
	GenerationRate time.Duration `env:"GENERATION_RATE" envDefault:"250ms"`

	ConsumerWarmup time.Duration `env:"CONSUMER_WARMUP" envDefault:"5m"`
}
