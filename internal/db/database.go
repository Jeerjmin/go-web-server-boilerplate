package database

import (
	"fmt"
	"go-auth/internal/config"
	"go-auth/internal/db/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	DB     *gorm.DB
	Config *config.Config
}

func NewDatabase(config *config.Config) *Database {
	return &Database{
		Config: config,
	}
}

func (db *Database) Connect() {
	dsn := "host=" + db.Config.PostgresqlHost + " port=" + db.Config.PostgresqlPort + " user=" + db.Config.PostgresqlUser + " password=" + db.Config.PostgresqlPassword + " dbname=" + db.Config.PostgresqlDatabase + " sslmode=disable"
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}

	sqlDB, err := conn.DB()
	if err != nil {
		log.Fatal("Error getting database connection: ", err)
	}

	if err := sqlDB.Ping(); err != nil {
		log.Fatal("Error pinging the database: ", err)
	}

	fmt.Println("Connected to the database")

	db.DB = conn

	db.AutoMigrate()
}

func (db *Database) AutoMigrate() {
	err := db.DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Error during migration: ", err)
	}
	fmt.Println("Database migration completed")
}
