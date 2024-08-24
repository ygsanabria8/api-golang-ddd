package server

import (
	"github.com/spf13/viper"
	"strings"
)

func SetUpConfig(logger *Logger) *Configuration {
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		logger.Errorw("Error Reading .Env File", err)
		panic(err.Error())
	}
	viper.AutomaticEnv()

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.MergeInConfig(); err != nil {
		logger.Errorw("Error Reading config.yaml", err)
		panic(err.Error())
	}

	var config Configuration
	if err := viper.Unmarshal(&config); err != nil {
		logger.Errorw("Error Decoding Config Object", err)
		panic(err.Error())
	}

	config.Server.Header = strings.Split(viper.GetString("Header"), ",")
	config.Server.Origin = strings.Split(viper.GetString("Origin"), ",")
	return &config
}
