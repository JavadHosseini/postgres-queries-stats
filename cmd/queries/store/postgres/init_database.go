package postgres

import (
	"fmt"
	"os"
	"time"

	"gorm.io/driver/postgres"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Database
type Database struct {
	*gorm.DB
}

var DB *gorm.DB

// Init This function is responsible for
// initializing new database connection
func Init() (*gorm.DB, error) {
	//Get database details from environment variables
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	DBName := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		host, port, user, DBName, password,
	)
	dbConn, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, _ := dbConn.Session(&gorm.Session{FullSaveAssociations: true}).DB()
	sqlDB.SetMaxIdleConns(3)
	sqlDB.SetMaxOpenConns(1000)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)
	// sqlDB.LogMode(true)

	DB = dbConn
	return DB, nil

}
