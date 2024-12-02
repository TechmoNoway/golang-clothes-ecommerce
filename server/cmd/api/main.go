package main

import (
	"log"

	"github.com/TechmoNoway/golang-clothes-ecommerce/internal/env"
	"github.com/TechmoNoway/golang-clothes-ecommerce/internal/store"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	store := store.NewStorage(nil)

	cfg := config{
		addr: env.GetString("ADDR", ":4243"),
	}
	app := &application{
		config: cfg,
		store:  store,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))

}
