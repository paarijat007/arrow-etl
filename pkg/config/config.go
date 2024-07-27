package config

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
)

// Config represents the configuration structure
type Config struct {
	LogLevel               string `json:"log_level"`
	MaxConcurrentJobs      int    `json:"max_concurrent_jobs"`
	EnableParallel         bool   `json:"enable_parallel"`
	ChunkSize              int    `json:"chunk_size"`
	DefaultSourceType      string `json:"default_source_type"`
	DefaultDestinationType string `json:"default_destination_type"`
}

var (
	config *Config
	once   sync.Once
)

// GetConfig returns the singleton config instance
func GetConfig() *Config {
	once.Do(func() {
		config = &Config{
			LogLevel:               "INFO",
			MaxConcurrentJobs:      5,
			EnableParallel:         true,
			ChunkSize:              1000,
			DefaultSourceType:      "MySQL",
			DefaultDestinationType: "CSV",
		}
	})
	return config
}

// LoadConfig loads the configuration from a JSON file
func LoadConfig(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("error opening config file: %v", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	config := GetConfig()
	if err := decoder.Decode(config); err != nil {
		return fmt.Errorf("error decoding config file: %v", err)
	}

	return nil
}

// SaveConfig saves the current configuration to a JSON file
func SaveConfig(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("error creating config file: %v", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(GetConfig()); err != nil {
		return fmt.Errorf("error encoding config: %v", err)
	}

	return nil
}

// SetLogLevel sets the log level
func SetLogLevel(level string) {
	GetConfig().LogLevel = level
}

// SetMaxConcurrentJobs sets the maximum number of concurrent jobs
func SetMaxConcurrentJobs(max int) {
	GetConfig().MaxConcurrentJobs = max
}

// SetEnableParallel sets whether parallel processing is enabled
func SetEnableParallel(enable bool) {
	GetConfig().EnableParallel = enable
}

// SetChunkSize sets the chunk size for data processing
func SetChunkSize(size int) {
	GetConfig().ChunkSize = size
}

// SetDefaultSourceType sets the default source type
func SetDefaultSourceType(sourceType string) {
	GetConfig().DefaultSourceType = sourceType
}

// SetDefaultDestinationType sets the default destination type
func SetDefaultDestinationType(destType string) {
	GetConfig().DefaultDestinationType = destType
}
