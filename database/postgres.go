package database

import (
	"fmt"
	"strconv"

	"github.com/MarcelArt/app_standard/config"
	"github.com/MarcelArt/app_standard/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func ConnectDB() {
	p := config.Env.DBPort
	port, err := strconv.ParseUint(p, 10, 32)

	if err != nil {
		panic("failed to parse database port")
	}

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Env.DBHost, port, config.Env.DBUser, config.Env.DBPassword, config.Env.DBName)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connection Opened to Database")

	// MigrateDB()
}

func GetDB() *gorm.DB {
	return db
}

func MigrateDB() {
	db.AutoMigrate(
		models.Template{},
		models.Process{},
		models.User{},
		models.AuthorizedDevice{},
	)
	fmt.Println("Database Migrated")
}

func DropDB() {
	db.Migrator().DropTable(
		&models.Template{},
		&models.Process{},
		models.User{},
		models.AuthorizedDevice{},
	)
	fmt.Println("Database Droped")
}
