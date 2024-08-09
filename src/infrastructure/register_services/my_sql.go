package register_services

import (
	"api.ddd/src/api/server"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ProvideSqlClient Provide Sql Client Instance
func ProvideSqlClient(logger *server.Logger) *gorm.DB {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)?charset=utf8&parseTime=True&loc=Local",
		"admin",
		"123456",
		"127.0.0.1",
		3306,
	)
	client, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Fatalf("Cannot Connect With Sql %v", err.Error())
		panic(err.Error())
	}

	db, _ := client.DB()
	db.SetMaxIdleConns(50)
	db.SetMaxOpenConns(50)

	return client
}
