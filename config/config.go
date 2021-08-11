package config

type Сonfig struct {
	RedisPass string `env:"REDIS_PASS" envDefault:""`
	RedisHost string `env:"REDIS_HOST" envDefault:"localhost"`
	RedisPort int    `env:"REDIS_PORT" envDefault:"6379"`
}
