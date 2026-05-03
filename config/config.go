package config

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	App      AppConfig      `mapstructure:"app"`
	Database DatabaseConfig `mapstructure:"database"`
}

type AppConfig struct {
	Port         string `mapstructure:"port"`
	Env          string `mapstructure:"env"`
	DbConnection string `mapstructure:"dbConnection"`
}

type DatabaseConfig struct {
	Host            string        `mapstructure:"host"`
	Port            string        `mapstructure:"port"`
	User            string        `mapstructure:"user"`
	Password        string        `mapstructure:"password"`
	DBName          string        `mapstructure:"dbName"`
	SSLMode         string        `mapstructure:"sslMode"`
	MaxIdleConns    int           `mapstructure:"maxIdleConns"`
	MaxOpenConns    int           `mapstructure:"maxOpenConns"`
	ConnMaxLifeTime time.Duration `mapstructure:"connMaxLifeTime"`
}

func LoadConfig(configDir string) (*Config, error) {
	env := os.Getenv("APP_ENV")

	if env == "" {
		env = "development"
	}

	configName := fmt.Sprintf("config-%s", env)

	viper.AddConfigPath(configDir)
	viper.SetConfigName(configName)
	viper.SetConfigType("yml")

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config file %v", err)
	}

	var cfg Config

	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("failed to load config %v", err)
	}

	return &cfg, nil
}
