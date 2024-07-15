package main

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type MongoConfig struct {
	Host    string `yaml:"host"`
	Port    int    `yaml:"port"`
	DB      string `yaml:"db"`
	User    string `yaml:"user"`
	Pass    string `yaml:"pass"`
	Collect string `yaml:"collect"`
}

func (c *MongoConfig) getConf() *MongoConfig {
	yamlFile, err := os.ReadFile("config.yaml")
	if err != nil {
		log.Printf("Failed getting config")
		return nil
	}

	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Printf("Failed unmarshalling config")
		return nil
	}

	return c
}
