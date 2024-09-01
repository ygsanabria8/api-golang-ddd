package register_services

import (
	"api.ddd/src/api/server"
	"api.ddd/src/domain/aggregates"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ProvideSqlClient Provide Sql Client Instance
func ProvideSqlClient(logger *server.Logger, config *server.Configuration) *gorm.DB {
	dsn := fmt.Sprintf(
		"%s/%s?charset=utf8&parseTime=True&loc=Local",
		config.Sql.ConnectionString,
		config.Sql.DatabaseName,
	)
	client, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Fatalf("Cannot Connect With Sql %v", err.Error())
		panic(err.Error())
	}

	db, _ := client.DB()
	db.SetMaxIdleConns(50)
	db.SetMaxOpenConns(50)

	err = client.AutoMigrate(&aggregates.User{})
	if err != nil {
		logger.Fatalf("Cannot Migrate Into Sql %v", err.Error())
		panic(err.Error())
	}

	return client
}
