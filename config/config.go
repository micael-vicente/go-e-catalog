package config

// Config is the struct that holds the configuration for the application
type Config struct {
	Server  ServerConfig  `yaml:"server"`
	Elastic ElasticConfig `yaml:"elastic"`
}

// ElasticConfig is the struct that holds the configuration for Elasticsearch
type ElasticConfig struct {
	Url      string `yaml:"url"`
	Index    string `yaml:"index"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

// ServerConfig is the struct that holds the configuration for the server
type ServerConfig struct {
	Port int `yaml:"port"`
}
