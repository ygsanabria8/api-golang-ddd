package server

import (
	"github.com/spf13/viper"
	"os"
	"strings"
)

func SetUpConfig(logger *Logger) *Configuration {
	// In local viper load .env file
	if _, err := os.Stat(".env"); !os.IsNotExist(err) {
		viper.SetConfigFile(".env")
		if err := viper.ReadInConfig(); err != nil {
			logger.Warnf("Error reading .env file: %v", err)
		}
	}

	// Viper read environment variables
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

	config.Server.Header = strings.Split(viper.GetString("HEADER"), ",")
	config.Server.Origin = strings.Split(viper.GetString("ORIGIN"), ",")
	config.Mongo.ConnectionString = viper.GetString("MONGO_CONNECTION_STRING")
	config.Sql.ConnectionString = viper.GetString("SQL_CONNECTION_STRING")
	config.Kafka.Broker = viper.GetString("KAFKA_BROKER")

	return &config
}
