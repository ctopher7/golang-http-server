package config

import (
	"fmt"
	"os"
	"path/filepath"

	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	ServerAddress string     `yaml:"server_address"`
	PostgresDb    PostgresDb `yaml:"postgres_db"`
}

type PostgresDb struct {
	Host     string
	Port     int
	User     string
	Password string
	Dbname   string
}

func ReadConfig(env string) (*Config, error) {
	config := &Config{}

	absPath, _ := filepath.Abs(fmt.Sprintf("./berat/files/%s.yaml", env))
	file, err := os.Open(absPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	d := yaml.NewDecoder(file)

	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}
