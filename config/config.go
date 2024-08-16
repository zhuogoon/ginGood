package config

import (
	"errors"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Database struct {
		DSN string `yaml:"dsn"`
	} `yaml:"database"`
}

var AppConfig Config

func LoadConfig() error {
	file, err := os.Open("config/config.yaml")
	if err != nil {
		return errors.New("没找打yaml文件")
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&AppConfig); err != nil {
		return err
	}
	return nil
}
