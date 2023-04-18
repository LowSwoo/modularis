package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Server struct {
		Host string `yaml:"Host"`
		Port string `yaml:"Port"`
	} `yaml:"Server"`
	MongoDatabase struct {
		Host           string `yaml:"Host"`
		Port           string `yaml:"Port"`
		DbName         string `yaml:"DbName"`
		CollectionName string `yaml:"CollectionName"`
	} `yaml:"Mongo"`
}

func NewConfig(filePath string) *Config {
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	var cfg Config
	if err = yaml.NewDecoder(f).Decode(&cfg); err != nil {
		panic(err)
	}
	return &cfg
}
