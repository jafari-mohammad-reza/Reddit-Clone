package config

import (
	"log"
	"os"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

func GetConfig() *Config {

	filePath := getConfigPath(os.Getenv("APP_ENV"))
	v, loadErr := LoadConfigFile(filePath)
	if loadErr != nil {
		log.Fatalf("Unable to load config: %v\n", loadErr)
	}
	cfg, parseErr := ParseConfig(v)
	if parseErr != nil {
		log.Fatalf("Unable to parse config: %v\n", parseErr)
	}
	return cfg
}

func ParseConfig(v *viper.Viper) (*Config, error) {
	var cfg Config
	err := v.Unmarshal(&cfg)
	if err != nil {
		log.Printf("Unable to read config: %v", err)
		return nil, err
	}
	return &cfg, nil
}

func LoadConfigFile(filePath string) (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigType("yaml")
	v.SetConfigName(filePath)
	v.AddConfigPath(".")
	v.AutomaticEnv()
	v.Debug()

	err := v.ReadInConfig()
	if err != nil {
		log.Printf("Unable to read config: %v", err)
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			return nil, errors.New("config file not found")
		}
		return nil, err
	}
	return v, nil

}

func getConfigPath(env string) string {
	switch env {
	case "development":
		return "/src/share/config/config-development.yml"
	case "production":
		return "/src/share/config/config-development.yml"
	case "docker":
		return "/src/share/config/config-development.yml"
	default:
		return "/src/share/config/config-development.yml"
	}
}
