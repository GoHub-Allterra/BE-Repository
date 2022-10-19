package database

import (
	// "gohub/config"
	posts "gohub/features/post/repository"
	user "gohub/features/user/repository"
	"os"

	"github.com/labstack/gommon/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {


	str := os.Getenv("DB_USER")+":"+os.Getenv("DB_PWD")+"@tcp("+os.Getenv("DB_HOST")+":"+os.Getenv("DB_PORT")+")/"+os.Getenv("DB_NAME")+"?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(str), &gorm.Config{})
	if err != nil {
		log.Error("db config error :", err.Error())
		return nil
	}
	migrateDB(db)
	return db
}

func migrateDB(db *gorm.DB) {
	db.AutoMigrate(&user.User{})
	db.AutoMigrate(&posts.Post{})
}
