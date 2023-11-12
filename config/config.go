package config

import (
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	App struct {
		Host    string        `mapstructure:"HOST"`
		Port    string        `mapstructure:"PORT"`
		Timeout time.Duration `mapstructure:"TIMEOUT"`
	} `mapstructure:"APP"`

	Mysql struct {
		Host     string `mapstructure:"HOST"`
		Port     string `mapstructure:"PORT"`
		Database string `mapstructure:"DATABASE"`
		Username string `mapstructure:"USERNAME"`
		Password string `mapstructure:"PASSWORD"`
	} `mapstructure:"MYSQL"`
}

func Get() *Config {
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		logrus.Errorf("Viper ReadIn: %s", err)
		return nil
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		logrus.Errorf("viper unmarshal: %s", err)
		return nil
	}

	return &config
}
