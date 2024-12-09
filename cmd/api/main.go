package main

import (
	"log"

	"github.com/TechmoNoway/golang-clothes-ecommerce/internal/db"
	"github.com/TechmoNoway/golang-clothes-ecommerce/internal/env"
	"github.com/TechmoNoway/golang-clothes-ecommerce/internal/store"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	cfg := config{
		addr: env.GetString("ADDR", ":4243"),
		db: dbConfig{
			addr:         env.GetString("DB_ADDR", "postgres://admin:adminpassword@localhost/clothesecommerce?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
	}

	db, err := db.New(
		cfg.db.addr,
		cfg.db.maxOpenConns,
		cfg.db.maxIdleConns,
		cfg.db.maxIdleTime,
	)

	store := store.NewStorage(db)

	if err != nil {
		log.Panic(err)
	}

	defer db.Close()
	log.Print("database connection pool established")

	app := &application{
		config: cfg,
		store:  store,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))

}
