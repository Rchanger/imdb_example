package models

import (
	"log"
	"os"
)

var (
	// Config - Global Configuration
	Config     AppConfig
	outfile, _ = os.Create("imdb.log")
	// Logger - Logger
	Logger = log.New(outfile, "IMDB:", log.Ldate|log.Ltime|log.Llongfile)
)

// GetConfigFilePath return config filepath
func GetConfigFilePath() string {
	return "config/config.toml"
}
