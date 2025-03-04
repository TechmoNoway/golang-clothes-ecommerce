package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/TechmoNoway/golang-clothes-ecommerce/internal/auth"
	"github.com/TechmoNoway/golang-clothes-ecommerce/internal/store"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"go.uber.org/zap"
)

type application struct {
	config        config
	store         store.Storage
	logger        *zap.SugaredLogger
	authenticator auth.Authenticator
}

type config struct {
	addr string
	db   dbConfig
	auth authConfig
}

type dbConfig struct {
	addr         string
	maxOpenConns int
	maxIdleConns int
	maxIdleTime  string
}

type authConfig struct {
	basic basicConfig
	token tokenConfig
}

type tokenConfig struct {
	secret string
	exp    time.Duration
	iss    string
}

type basicConfig struct {
	user string
	pass string
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
			r.Get("/health", app.healthCheckHandler)

			r.Route("/users", func(r chi.Router) {
				r.Route("/getUserByID/{userID}", func(r chi.Router) {
					r.Get("/", app.getUserHandler)
				})
				r.Get("/getAllUsers", app.getAllUsersHandler)
				r.Put("/updateUser", app.updateUserHandler)
			})

			r.Route("/roles", func(r chi.Router) {
				r.Get("/", func(w http.ResponseWriter, r *http.Request) {
					fmt.Fprintln(w, "This is roles api")
				})
			})

			r.Route("/products", func(r chi.Router) {
				r.Get("/getAllProducts", app.getAllProductsHandler)
				r.Get("/getAllProductsByName", app.getAllProductsByNameHandler)
				r.Get("/getAllProductsByCategoryID", app.getAllProductsByCategoryIDHandler)
				r.Post("/createProduct", app.createProductHandler)
				r.Put("/updateProduct", app.updateProductHandler)
				r.Delete("/deleteProduct", app.deleteProductHandler)
			})

			r.Route("/orders", func(r chi.Router) {
				r.Post("/createOrder", app.createOrderHandler)
				r.Get("/getAllOrders", app.getAllOrdersHandler)
				r.Get("/getAllOrdersByUserID", app.getAllOrdersByUserIDHandler)
				r.Put("/updateOrder", app.updateOrderHandler)
			})

			r.Route("/orderitems", func(r chi.Router) {
				r.Get("/getAllOrderItems", app.getAllOrdersHandler)
				r.Post("/createOrderItem", app.createOrderItemHandler)
				r.Get("/getAllOrderItemsByOrderID", app.getAllOrderItemsByOrderIDHandler)
			})

			r.Route("/categories", func(r chi.Router) {
				r.Get("/getAllCategories", app.getAllCategoriesHandler)
				r.Post("/createCategory", app.createCategoryHandler)
			})

			r.Route("/auth", func(r chi.Router) {
				r.Post("/register", app.registerUserHandler)
				r.Post("/login", app.loginUserHandler)
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
