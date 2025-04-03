package config

import (
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type Config struct {
	App struct {
		Name  string `mapstructure:"name"`
		Debug bool   `mapstructure:"debug"`
	} `mapstructure:"app"`
	
	Features struct {
		Obsidian struct {
			VaultPath string `mapstructure:"vault_path"`
		} `mapstructure:"obsidian"`
		Placeholder struct {
			DefaultFile string `mapstructure:"default_file"`
		} `mapstructure:"placeholder"`
	} `mapstructure:"features"`
	
	OpenAIAPIKey string
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	
	// Add home directory config path
	home, err := os.UserHomeDir()
	if err == nil {
		viper.AddConfigPath(filepath.Join(home, ".config/cli"))
	}
	
	// Bind environment variables
	viper.AutomaticEnv()
	viper.BindEnv("OpenAIAPIKey", "OPENAI_API_KEY")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
} 