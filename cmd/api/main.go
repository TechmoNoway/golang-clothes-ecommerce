package main

import (
	"log"
	"time"

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
			addr:         env.GetString("DB_ADDR", "postgres://postgres:332003@localhost/clothesecommerce?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
		auth: authConfig{
			basic: basicConfig{
				user: env.GetString("AUTH_BASIC_USER", "admin"),
				pass: env.GetString("AUTH_BASIC_PASS", "admin"),
			},
			token: tokenConfig{
				secret: env.GetString("AUTH_TOKEN_SECRET", "example"),
				exp:    time.Hour * 24 * 3,
				iss:    "golangclothesecommerce",
			},
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
