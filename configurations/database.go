package configurations

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var DB *sql.DB
var DBgorm *gorm.DB

func InitDatabase() {
	db, err := sql.Open("sqlserver", "server=localhost;port=1500;user id=sa;password=docker2018!; database=Budget")
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}
	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("Connected!\n")
	DB = db
}

func InitDatabaseGorm() {
	dsn := "tcp://localhost:1500?database=Budget&username=sa&password=docker2018!&read_timeout=10&write_timeout=20"
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}

	DBgorm = db
}
