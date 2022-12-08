package main

import (
	"log"

	paths "boris-middleware/paths"

	"github.com/gofiber/fiber/v2"

	"database/sql"

	"github.com/go-sql-driver/mysql"
)

// Database instance
var db *sql.DB

// Connect function
func Connect() error {
	cfg := mysql.Config{
        User:   "root",
        Passwd: "0o0p0o0p",
        Net:    "tcp",
        Addr:   "127.0.0.1:3306",
        DBName: "boris_recipes",
    }

	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return err
	}
	if err = db.Ping(); err != nil {
		return err
	}
	return nil
}


func main(){

    if err := Connect(); err != nil {
		log.Fatal(err)
	}

	method_resolver := paths.CookingMethodTable(db)

    app := fiber.New()

    app.Post("/recipes", method_resolver)

    log.Fatal(app.Listen(":4000"))
}