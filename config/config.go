package config

import (
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	DBDriver string `mapstructure:"DB_DRIVER"`
	DBUrl    string `mapstructure:"DB_URL"`
	PORT     string `mapstructure:"PORT"`
}

func LoadConfig() (c Config, err error) {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		viper.SetDefault("DB_DRIVER", "postgres")
		viper.SetDefault("DB_URL", "localhost:5432")

		for _, key := range []string{
			"DB_DRIVER",
			"DB_URL",
			"PORT",
		} {
			envVal := os.Getenv(key)
			if envVal != "" {
				viper.Set(key, envVal)
			}
		}
	}

	err = viper.Unmarshal(&c)
	return
}
