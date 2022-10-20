package database

import (
	// "gohub/config"
	"os"
	posts "gohub/features/post/repository"
	user "gohub/features/user/repository"


	"github.com/labstack/gommon/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {


	str := os.Getenv("DB_USER")+":"+os.Getenv("DB_PWD")+"@tcp("+os.Getenv("DB_HOST")+":"+os.Getenv("DB_PORT")+")/"+os.Getenv("DB_NAME")+"?charset=utf8mb4&parseTime=True&loc=Local"

	// str := fmt.Sprint("root:@tcp(127.0.0.1:3306)/gohub?charset=utf8mb4&parseTime=True&loc=Local")
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
