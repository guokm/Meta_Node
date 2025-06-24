package config

import (
	"time"

	"github.com/spf13/viper"
)

var (
	GlobalConfig *BaseConfig
	configPath   = "config.yaml"
)

type BaseConfig struct {
	Name  string       `mapstructure:"Name"`
	Env   string       `mapstructure:"Env"`
	Port  string       `mapstructure:"Port"`
	Log   LoggerConfig `mapstructure:"Log"`
	Mysql MysqlConfig  `mapstructure:"Mysql"`
	Jwt   JwtConfig    `mapstructure:"Jwt"`
}

type (
	MysqlConfig struct {
		UserName        string `mapstructure:"Username"`
		Password        string `mapstructure:"Password"`
		Host            string `mapstructure:"Host"`
		Port            string `mapstructure:"Port"`
		DBName          string `mapstructure:"DBName"`
		Timeout         string `mapstructure:"Timeout"`
		DSN             string `mapstructure:"DSN"`
		MaxOpenConns    int    `mapstructure:"MaxOpenConns"`
		MaxIdleConns    int    `mapstructure:"MaxIdleConns"`
		ConnMaxLifetime string `mapstructure:"ConnMaxLifetime"`
	}

	LoggerConfig struct {
		Level      string `mapstructure:"Level"`
		LogPath    string `mapstructure:"LogPath"`
		Model      string `mapstructure:"Model"`
		MaxSize    int    `mapstructure:"MaxSize"`
		MaxBackups int    `mapstructure:"MaxBackups"`
		MaxAge     int    `mapstructure:"MaxAge"`
		Compress   bool   `mapstructure:"Compress"`
	}

	JwtConfig struct {
		Secret      string        `mapstructure:"Secret"`
		TokenExpire time.Duration `mapstructure:"TokenExpire"`
		Issuer      string        `mapstructure:"Issuer"`
	}
)

func LoadConfig() error {
	GlobalConfig = &BaseConfig{}
	viper.SetConfigFile(configPath)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	// 读配置
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	if err := viper.Unmarshal(&GlobalConfig); err != nil {
		return err
	}
	return nil
}
