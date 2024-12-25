package main

import (
	"log"

	"github.com/TechmoNoway/golang-clothes-ecommerce/internal/db"
	"github.com/TechmoNoway/golang-clothes-ecommerce/internal/env"
	"github.com/TechmoNoway/golang-clothes-ecommerce/internal/store"
)

func main() {
	addr := env.GetString("DB_ADDR", "postgres://postgres:332003@localhost/clothesecommerce?sslmode=disable")
	conn, err := db.New(addr, 3, 3, "15m")
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	store := store.NewStorage(conn)

	db.Seed(store, conn)
}
