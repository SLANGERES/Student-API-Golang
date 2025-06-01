package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HttpConfig struct {
	Addr string `yaml:"address" env-required:"true"`
}

// Setting up the configuration file
type Config struct {
	Env         string `yaml:"env" env:"ENV" env-required:"true" env-default:"production"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HttpConfig  `yaml:"http_server"`
}

//Config function that must be run first

func MustDone() *Config {
	var configPath string

	// First check environment variable
	configPath = os.Getenv("CONFIG_PATH")

	// Then fallback to command-line flag
	if configPath == "" {
		flagPath := flag.String("config", "", "Path to the configuration file")
		flag.Parse()
		configPath = *flagPath

		if configPath == "" {
			log.Fatal("Configuration file path must be provided via CONFIG_PATH env variable or --config flag")
		}
	}

	// Check if config file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("Config file does not exist: %s", configPath)
	}

	var cfg Config
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("Unable to parse config data: %v", err)
	}

	return &cfg
}
