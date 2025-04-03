package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type Config struct {
	App struct {
		Name  string `mapstructure:"name"`
		Debug bool   `mapstructure:"debug"`
	} `mapstructure:"app"`
	
	ObsidianVaultPath string `mapstructure:"obsidian_vault_path"`
	OpenAIAPIKey string `mapstructure:"openai_api_key"`
	AnthropicAPIKey string `mapstructure:"anthropic_api_key"`
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
	viper.BindEnv("openai_api_key", "OPENAI_API_KEY")
	viper.BindEnv("anthropic_api_key", "ANTHROPIC_API_KEY")
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	if config.OpenAIAPIKey == "" && config.AnthropicAPIKey == "" {
		return nil, fmt.Errorf("OPENAI_API_KEY or ANTHROPIC_API_KEY environment variable is not set")
	}

	if config.ObsidianVaultPath == "" {
		return nil, fmt.Errorf("OBSIDIAN_VAULT_PATH environment variable is not set")
	}


	return &config, nil
} 