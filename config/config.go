package config

import (
	"be21/users"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// InitMysql initializes the MySQL database connection and returns a pointer to a gorm.DB object.
func InitMysql() *gorm.DB {
	connectionString := "host=localhost user=postgres password=TEAMSECRETGG dbname=crud_barang port=6666 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
	if err != nil {
		fmt.Println("terjadi sebuah kesalahan", err.Error())
		return nil
	}
	return db
}

// Migrate performs database migration for the TopUps, Transfers, and Users tables.
func Migrate(connection *gorm.DB) error {
	err := connection.AutoMigrate(&users.TopUp{})
	if err != nil {
		return err
	}
	err = connection.AutoMigrate(&users.Transfer{})
	if err != nil {
		return err
	}
	err = connection.AutoMigrate(&users.User{})
	return err
}
