package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"lion-test/helper"
)

func NewDB() *gorm.DB {
	dbUser := helper.GoDotEnvVariable("DB_USER")
	dbPass := helper.GoDotEnvVariable("DB_PASS")
	dbName := helper.GoDotEnvVariable("DB_NAME")
	dbHost := helper.GoDotEnvVariable("DB_HOST")
	dbPort := helper.GoDotEnvVariable("DB_PORT")
	dbDsn := SetDB(dbName, dbUser, dbPass, dbHost, dbPort)
	db, err := gorm.Open(mysql.Open(dbDsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	helper.PanicIfError(err)

	return db
}

func SetDB(db_name, db_user, db_pass, db_host, db_port string) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", db_user, db_pass, db_host, db_port, db_name)
}
