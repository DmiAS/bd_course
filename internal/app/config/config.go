package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type HTTPConfig struct {
	Port string
}
type DBConfig struct {
	AdminDSN  string
	WorkerDSN string
	ClientDSN string
	DummyDSN  string
}

type Config struct {
	HTTP HTTPConfig
	DB   DBConfig
}

const (
	ConfigName = "ConfigName"
	ConfigPath = "ConfigPath"
	ConfigType = "ConfigType"
)

func NewConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}
	configName := os.Getenv(ConfigName)
	configPath := os.Getenv(ConfigPath)
	configType := os.Getenv(ConfigType)
	viper.SetConfigName(configName) // name of config file (without extension)
	viper.SetConfigType(configType) // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(configPath) // call multiple times to add many search paths
	viper.AddConfigPath(".")        // optionally look for config in the working directory
	err := viper.ReadInConfig()     // Find and read the config file
	if err != nil {                 // Handle errors reading the config file
		return nil, err
	}

	httpCfg := HTTPConfig{}
	if err := viper.Unmarshal(&httpCfg); err != nil {
		return nil, err
	}
	dbCfg := DBConfig{}
	if err := viper.Unmarshal(&dbCfg); err != nil {
		return nil, err
	}
	return &Config{
		HTTP: httpCfg,
		DB:   dbCfg,
	}, nil
}
