package db

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/yaska1706/AfyaPevu/internal/pkg/config"
	"github.com/yaska1706/AfyaPevu/internal/pkg/models/tasks"
	"github.com/yaska1706/AfyaPevu/internal/pkg/models/users"
)

var (
	DB  *gorm.DB
	err error
)

type Database struct {
	*gorm.DB
}

// SetupDB opens a database and saves the reference to `Database` struct.
func SetupDB() {
	var db = DB

	configuration := config.GetConfig()

	database := configuration.Database.Dbname
	username := configuration.Database.Username
	password := configuration.Database.Password
	host := configuration.Database.Host
	port := configuration.Database.Port

	db, err = gorm.Open("postgres", "host="+host+" port="+port+" user="+username+" dbname="+database+"  sslmode=disable password="+password)
	if err != nil {
		fmt.Println("db err: ", err)
	}

	// Change this to true if you want to see SQL queries
	db.LogMode(true)
	db.DB().SetMaxIdleConns(configuration.Database.MaxIdleConns)
	db.DB().SetMaxOpenConns(configuration.Database.MaxOpenConns)
	db.DB().SetConnMaxLifetime(time.Duration(configuration.Database.MaxLifetime) * time.Second)
	DB = db
	migration()
}

// Auto migrate project models
func migration() {
	DB.AutoMigrate(&users.User{})
	DB.AutoMigrate(&users.UserRole{})
	DB.AutoMigrate(&tasks.Task{})
}

func GetDB() *gorm.DB {
	return DB
}
