package config

import (
	"time"
	"os"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"github.com/joho/godotenv"
)
type Config struct {
	Env string `yaml:"env" env:"ENV" env-default:"local"` // env-required:"true" - сделатть обязательным 
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HTTTPServer `yaml:"http-server"`
}

type HTTTPServer struct {
	Address string `yaml:"address" env-default:"localhost:8080"`
	Timeout time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTtimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

//Must принято что может паниковать а не возвращить ошибку
func MustLoad() *Config {
	err := godotenv.Load("../../local.env ")
	if err != nil {
		log.Fatalf("Error when loading config: %s", err)
	}
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exsts: %s", configPath)
	}
	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatal("Cannot read config: %s", err)
	}
	return &cfg
}