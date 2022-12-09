package main

import (
	"fmt"
	"log"

	paths "boris-middleware/paths"

	"github.com/gofiber/fiber/v2"

	"database/sql"

	"github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2/middleware/cors"
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

    app := fiber.New()
	app.Use(cors.New())
	app.Get("/", func(ctx *fiber.Ctx) error{
		return ctx.SendString("alive")
	})

    app.Post("/laxrecipe", func(ctx *fiber.Ctx) error{
		return paths.CookingMethodLax(ctx, db)
	})
    app.Post("/strictrecipe", func(ctx *fiber.Ctx) error{
		return paths.CookingMethodStrict(ctx, db)
	})
	app.Static("/media", "./media")

	fmt.Println("starting fiber")
    log.Fatal(app.Listen(":4000"))
}