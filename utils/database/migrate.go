package database

import (
	"fmt"

	// DataFashion "kos-barokah-api/features/fashions/data"
	DataRoom "kos-barokah-api/features/room/data"
	DataUser "kos-barokah-api/features/users/data"

	// DataVoucher "kos-barokah-api/features/vouchers/data"

	"gorm.io/gorm"
)

func MigrateWithDrop(db *gorm.DB) {
	// Drop entire schema
	// db.Exec("DROP DATABASE IF EXISTS defaultdb")
	// db.Exec("CREATE DATABASE defaultdb")

	// db.Exec("USE defaultdb")
	// fmt.Println("[MIGRATION] Success dropping aifash database and creating a new one")

	// USER DATA MANAGEMENT \\
	db.AutoMigrate(DataUser.User{})
	db.AutoMigrate(DataUser.UserResetPass{})
	fmt.Println("[MIGRATION] Success creating user")

	db.AutoMigrate(DataRoom.Room{})

}
