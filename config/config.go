package config

import "log"

// LoadConfig simply logs that we're using system environment variables
func LoadConfig() error {
	log.Println("Using system environment variables for configuration.")
	return nil
}
