package config

import (
	"path/filepath"
	"strings"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	App struct {
		Addr string `mapstructure:"addr"`
	}
	Redis struct {
		Addr         string        `mapstructure:"addr"`
		Password     string        `mapstructure:"password"`
		DB           int           `mapstructure:"db"`
		WriteTimeout time.Duration `mapstructure:"write_timeout"`
		ReadTimeout  time.Duration `mapstructure:"read_timeout"`
	}
	Database struct {
		Driver string `mapstructure:"driver"`
		DSN    string `mapstructure:"dsn"`
	}
}

func New(path string) (*Config, error) {
	absPath, _ := filepath.Abs(path)

	base := filepath.Base(absPath)
	dir := filepath.Dir(absPath)
	cfgFileName := strings.Split(base, ".")
	cfgName, cfgEnv := cfgFileName[0], cfgFileName[1]

	viper.SetConfigType("yaml")
	viper.SetConfigName(strings.Join([]string{cfgName, cfgEnv}, "."))
	viper.AddConfigPath(dir)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var c Config

	if err := viper.Unmarshal(&c); err != nil {
		return nil, err
	}

	return &c, nil
}
