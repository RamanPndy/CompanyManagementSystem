package config

import (
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
)

var configModel *Config

// NewConfig gets the configuration based on the environment passed
func NewConfig() (IConfig, error) {
	configFilePath, err := getConfigFilePath()
	if err != nil {
		return nil, err
	}
	bytes, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(bytes, &configModel)
	if err != nil {
		return nil, err
	}

	configModel.DB.Host = os.Getenv("DB_HOST")
	configModel.DB.Port = os.Getenv("DB_PORT")
	configModel.DB.User = os.Getenv("DB_USER")
	configModel.DB.Password = os.Getenv("DB_PWD")
	configModel.DB.Name = os.Getenv("DB_NAME")

	// Returns
	return &IConfigModel{model: configModel}, nil
}

func getConfigFilePath() (string, error) {
	env := os.Getenv("TIER")

	if env == "" {
		env = DEVELOPMENT
	}

	configFile := "/config/tier/" + env + ".yml"
	if env == DEVELOPMENT {
		dir, err := os.Getwd()
		if err != nil {
			return "", err
		}

		envFilePath := strings.ReplaceAll(dir, "/cmd/api", "/app.env")

		err = godotenv.Load(envFilePath)
		if err != nil {
			log.Fatal("Error loading .env file")
			return "", err
		}

		configFilePath := strings.ReplaceAll(dir, "/cmd/api", configFile)
		return configFilePath, nil
	}
	return strings.TrimPrefix(configFile, "/"), nil
}

// Get implements the interface function for IConfig
func (ic *IConfigModel) Get() *Config {
	return ic.model
}
