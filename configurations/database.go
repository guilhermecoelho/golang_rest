package configurations

import (
	"context"
	"database/sql"
	"fmt"
	"log"
)

var DB *sql.DB

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
