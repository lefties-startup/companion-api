package companion

import (
	"fmt"
	"github.com/lefties-startup/companion-api/connect/postgres"
	"github.com/spf13/viper"
)

type Config struct {
	Postgres PostgresManager `mapstructure:"postgres_manager"`
	HTTPAddr string          `mapstructure:"http_addr"`
}

type PostgresManager struct {
	Manager postgres.Config
}

func Load() (*Config, error) {
	// Указываем имя файла (без расширения) и путь
	viper.SetConfigName("companion.local.yaml")
	viper.SetConfigType("yaml")
	//viper.AddConfigPath(".")
	viper.AddConfigPath("./config")

	// Читаем конфиг
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return &cfg, nil
}
