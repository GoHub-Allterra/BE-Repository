// package database

// import (
// 	// "gohub/config"
// 	"os"
// 	// "fmt"
// 	comments "gohub/features/comments/repository"
// 	posts "gohub/features/post/repository"
// 	user "gohub/features/user/repository"

// 	"github.com/labstack/gommon/log"
// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// )

// func InitDB() *gorm.DB {

// 	str := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PWD") + "@tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/" + os.Getenv("DB_NAME") + "?charset=utf8mb4&parseTime=True&loc=Local"

// 	// str := fmt.Sprint("root:@tcp(127.0.0.1:3306)/gohub?charset=utf8mb4&parseTime=True&loc=Local")
// 	db, err := gorm.Open(mysql.Open(str), &gorm.Config{})
// 	if err != nil {
// 		log.Error("db config error :", err)
// 		return nil
// 	}
// 	migrateDB(db)
// 	return db
// }

// func migrateDB(db *gorm.DB) {
// 	db.AutoMigrate(&comments.Comments{})
// 	db.AutoMigrate(&user.User{})
// 	db.AutoMigrate(&posts.Post{})
// }

package database

import (
	// "gohub/config"
	// "os"
	"fmt"
	// "gohub/config"

	comments "gohub/features/comments/repository"
	posts "gohub/features/post/repository"
	user "gohub/features/user/repository"

	// "fmt"

	"github.com/labstack/gommon/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
// 	// str := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
// 	// 	c.DBUser,
// 	// 	c.DBPwd,
// 	// 	c.DBHost,
// 	// 	c.DBPort,
// 	// 	c.DBName,
// 	// )

	// str := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PWD") + "@tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/" + os.Getenv("DB_NAME") + "?charset=utf8mb4&parseTime=True&loc=Local"

	str := fmt.Sprint("root:@tcp(127.0.0.1:3306)/gohub?charset=utf8mb4&parseTime=True&loc=Local")
	// str := "root:@tcp(mysql:3306)/gohub?charset=utf8mb4&parseTime=True&loc=Local"

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
	db.AutoMigrate(&comments.Comments{})
}
