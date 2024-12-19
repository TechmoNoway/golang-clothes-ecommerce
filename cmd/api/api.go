package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/TechmoNoway/golang-clothes-ecommerce/internal/store"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"go.uber.org/zap"
)

type application struct {
	config config
	store  store.Storage
	logger *zap.SugaredLogger
}

type config struct {
	addr string
	db   dbConfig
}

type dbConfig struct {
	addr         string
	maxOpenConns int
	maxIdleConns int
	maxIdleTime  string
}

func (app *application) mount() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5174"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		ExposedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	r.Use(middleware.Timeout(60 * time.Second))

	r.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Get("/health", app.healthCheckHanler)

			r.Route("/users", func(r chi.Router) {
				r.Route("/getUserById/{userID}", func(r chi.Router) {
					r.Get("/", app.getUserHandler)
				})
				r.Get("/getAllUsers", app.getAllUsersHandler)

			})

			r.Route("/roles", func(r chi.Router) {
				r.Get("/", func(w http.ResponseWriter, r *http.Request) {
					fmt.Fprintln(w, "This is roles api")
				})
			})

			r.Route("/products", func(r chi.Router) {
				r.Get("/", func(w http.ResponseWriter, r *http.Request) {
					fmt.Fprintln(w, "This is products api")
				})
				r.Get("/getAllProducts", app.getAllProductsHanler)
				r.Route("/getAllProductsByName/{name}", func(r chi.Router) {
					r.Get("/", app.getAllProductsByNameHanler)
				})
				r.Route("/getAllProductsByCategoryID/{categoryID}", func(r chi.Router) {
					r.Get("/", app.getAllProductsByCategoryIDHanler)
				})
				r.Post("/createProduct", app.createProductHanler)
				r.Post("/updateProduct", app.updateProductHanler)
			})

			r.Route("/orders", func(r chi.Router) {
				r.Get("/", func(w http.ResponseWriter, r *http.Request) {
					fmt.Fprintln(w, "This is orders api")
				})
			})

			r.Route("/categories", func(r chi.Router) {

			})

			r.Route("/authentication", func(r chi.Router) {

			})

		})
	})

	return r
}

func (app *application) run(mux *chi.Mux) error {
	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	log.Printf("Server has started at %s", app.config.addr)

	return srv.ListenAndServe()
}
