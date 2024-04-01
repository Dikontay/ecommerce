package main

import (
	"ecommerce/internal/app"
	"ecommerce/internal/repository"
	"ecommerce/pkg"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := os.Getenv("DSN")
	db, err := pkg.OpenDB(dsn)
	if err != nil {
		log.Println(err)
		return
	}

	repo := repository.NewRepository(db)
	application := app.NewApp(repo)
	err = application.Start()
	if err != nil {
		log.Fatal(err)
	}
}
