package database

import (
	"gohub/config"
	user "gohub/features/user/repository"
	"github.com/labstack/gommon/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(c *config.AppConfig) *gorm.DB {
	// str := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
	// 	c.DBUser,
	// 	c.DBPwd,
	// 	c.DBHost,
	// 	c.DBPort,
	// 	c.DBName,
	// )

	str := "root:@tcp(mysql:3308)/gohub?charset=utf8mb4&parseTime=True&loc=Local"

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
}
