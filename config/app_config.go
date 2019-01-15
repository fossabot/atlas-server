package config

import "time"

// AppConfig holds the application configuration
type AppConfig struct {
	MongoConfig *MongoConfig `yaml:"mongo"`
}

// MongoConfig holds the mongo database configuration
type MongoConfig struct {
	Server         string        `yaml:"server"`
	Database       string        `yaml:"database"`
	ConnectTimeout time.Duration `yaml:"connect_timeout"`
}
