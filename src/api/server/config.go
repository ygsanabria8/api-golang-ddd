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

	allowedHeaders := viper.GetString("SERVER_HEADER")
	allowedOrigins := viper.GetString("SERVER_ORIGIN")

	headersSlice := strings.Split(allowedHeaders, ",")
	originsSlice := strings.Split(allowedOrigins, ",")

	return &Configuration{
		AppName: viper.GetString("appName"),
		Server: &Server{
			Port:   viper.GetString("server.port"),
			Origin: originsSlice,
			Header: headersSlice,
		},
		Mongo: &Mongo{
			Host:     viper.GetString("MONGO_HOST"),
			Port:     viper.GetString("MONGO_PORT"),
			User:     viper.GetString("MONGO_USER"),
			Password: viper.GetString("MONGO_PASSWORD"),
			Database: viper.GetString("mongo.database"),
			Collections: &Collections{
				User: viper.GetString("mongo.collections.user"),
			},
		},
		Sql: &Sql{
			Host:         viper.GetString("SQL_HOST"),
			Port:         viper.GetInt("SQL_PORT"),
			User:         viper.GetString("SQL_USER"),
			Password:     viper.GetString("SQL_PASSWORD"),
			DatabaseName: viper.GetString("sql.databaseName"),
		},
		Kafka: &Kafka{
			Broker: viper.GetString("KAFKA_BROKER"),
			Group:  viper.GetString("kafka.group"),
			Topics: &Topics{
				CreatedUser: viper.GetString("kafka.topics.createdUser"),
				UpdatedUser: viper.GetString("kafka.topics.updatedUser"),
				DeletedUser: viper.GetString("kafka.topics.deletedUser"),
			},
		},
	}
}
