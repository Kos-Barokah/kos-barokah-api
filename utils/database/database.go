package database

import (
	"fmt"

	"kos-barokah-api/configs"

	"github.com/labstack/gommon/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(c configs.ProgrammingConfig) (*gorm.DB, error) {

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", c.DBUser, c.DBPass, c.DBHost, c.DBPort, c.DBName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Error("Terjadi kesalahan pada database, error:", err.Error())
		return nil, err
	}

	return db, nil
}

// func Init() {

// 	DB.AutoMigrate(&models.User{})
// 	DB.AutoMigrate(&models.Blog{})
// 	DB.AutoMigrate(&models.Book{})
// 	DB.AutoMigrate(&models.Fashion{})
// 	DB.AutoMigrate(&models.Point{})
// 	DB.AutoMigrate(&models.Voucher{})
// 	DB.AutoMigrate(&models.UserVoucher{})

// }
