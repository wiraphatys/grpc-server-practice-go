package config

import (
	"log"
	"sync"

	"github.com/spf13/viper"
)

type Config struct {
	Server Server `mapstructure:",squash"`
	Db     Db     `mapstructure:",squash"`
	Jwt    Jwt    `mapstructure:",squash"`
}

type Server struct {
	Host string `mapstructure:"SERVER_HOST"`
	Port int    `mapstructure:"SERVER_PORT"`
}

type Db struct {
	Host     string `mapstructure:"DB_HOST"`
	Port     int    `mapstructure:"DB_PORT"`
	User     string `mapstructure:"DB_USER"`
	Password string `mapstructure:"DB_PASS"`
	Name     string `mapstructure:"DB_NAME"`
	SSLMode  string `mapstructure:"DB_SSL_MODE"`
	Timezone string `mapstructure:"DB_TIMEZONE"`
}

type Jwt struct {
	Secret     string `mapstructure:"JWT_SECRET"`
	Expiration int    `mapstructure:"JWT_EXPIRATION"`
}

type Grpc struct {
	UserUrl string `mapstructure:"GRPC_USER_URL"`
	AuthUrl string `mapstructure:"GRPC_AUTH_URL"`
}

var (
	once   sync.Once
	config *Config
)

func GetConfig() *Config {
	once.Do(func() {
		config = &Config{}

		viper.SetConfigFile("./.env")
		viper.AutomaticEnv()

		if err := viper.ReadInConfig(); err != nil {
			log.Fatalf("Error reading config file: %s", err)
		}

		if err := viper.Unmarshal(config); err != nil {
			log.Fatalf("Error unmarshalling config: %s", err)
		}
	})

	return config
}
