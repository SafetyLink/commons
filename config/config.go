package config

import (
	"github.com/spf13/viper"
	"os"
	"strings"
)

func ReadConfig[T any]() (*T, error) {
	viper.AddConfigPath("./")

	if os.Getenv("ENV") == "prod" {
		viper.SetConfigName("config.prod")
	} else {
		viper.SetConfigName("config.dev")
	}

	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	var cfg T
	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

func ReadConfigInTest[T any]() (*T, error) {
	workingDir, _ := os.Getwd()
	splitDir := strings.Split(workingDir, "\\")
	dirIndex := 0
	for i, v := range splitDir {
		if v == "internal" {
			dirIndex = i
			break
		}
	}
	viper.AddConfigPath(strings.Join(splitDir[:dirIndex], "\\"))

	if os.Getenv("ENV") == "prod" {
		viper.SetConfigName("config.prod")
	} else {
		viper.SetConfigName("config.dev")

	}

	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	var cfg T
	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
