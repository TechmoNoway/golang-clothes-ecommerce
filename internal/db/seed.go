package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/TechmoNoway/golang-clothes-ecommerce/internal/store"
)

var usernames = []string{
	"alice", "bob", "charlie", "dave", "eve", "frank", "grace", "heidi",
	"ivan", "judy", "karl", "laura", "mallory", "nina", "oscar", "peggy",
	"quinn", "rachel", "steve", "trent", "ursula", "victor", "wendy", "xander",
	"yvonne", "zack", "amber", "brian", "carol", "doug", "eric", "fiona",
	"george", "hannah", "ian", "jessica", "kevin", "lisa", "mike", "natalie",
	"oliver", "peter", "queen", "ron", "susan", "tim", "uma", "vicky",
	"walter", "xenia", "yasmin", "zoe",
}

var roleNames = []string{
	"user", "admin", "accounter",
}

var categoryNames = []string{
	"Men's Clothing", "Women's Clothing", "Kids' Clothing", "Accessories", "Footwear",
	"Sportswear", "Outerwear", "Underwear", "Formal Wear", "Casual Wear",
}

var productNames = []string{
	"Men's T-Shirt", "Women's Dress", "Kids' Jacket", "Men's Shoes", "Women's Handbag",
	"Sports Shorts", "Winter Coat", "Men's Underwear", "Business Suit", "Casual Jeans",
}

var productDescriptions = []string{
	"Comfortable cotton t-shirt", "Elegant evening dress", "Warm and cozy jacket", "Stylish leather shoes", "Fashionable handbag",
	"Breathable sports shorts", "Thick winter coat", "Soft and comfortable underwear", "Professional business suit", "Relaxed fit jeans",
}

var productStocks = []int64{
	100, 50, 75, 200, 150,
	120, 80, 90, 60, 110,
}

var productSize = []string{
	"S", "M", "L", "XL", "XXL",
}

var productColor = []string{
	"Red", "Blue", "Green", "Black", "White",
}

var productCategoryIds = []int64{
	1, 2, 3, 4, 5,
	6, 7, 8, 9, 10,
}

func Seed(store store.Storage, db *sql.DB) {
	ctx := context.Background()

	tables := []string{"users", "categories", "products", "roles"}
	for _, table := range tables {
		_, err := db.ExecContext(ctx, fmt.Sprintf("TRUNCATE TABLE %s RESTART IDENTITY CASCADE", table))
		if err != nil {
			log.Println("Error truncating table:", table, err)
			return
		}
	}

	roles := generateRoles(2)
	for _, role := range roles {
		err := store.Roles.Create(ctx, role)
		if err != nil {
			log.Println("Error creating role:", err)
			return
		}
	}

	users := generateUsers(5)

	for _, user := range users {
		err := store.Users.Create(ctx, user)
		if err != nil {
			log.Println("Error creating user:", err)
			return
		}
	}

	categories := generateCategories(10)
	for _, category := range categories {
		err := store.Categories.Create(ctx, category)
		if err != nil {
			log.Println("Error creating category:", err)
			return
		}
	}

	products := generateProducts(10)
	for _, product := range products {
		err := store.Products.Create(ctx, product)
		if err != nil {
			log.Println("Error creating product:", err)
			return
		}
	}

	log.Println("Seeding complete")
}

func generateUsers(num int) []*store.User {
	users := make([]*store.User, num)

	for i := 0; i < num; i++ {
		users[i] = &store.User{
			Username:  usernames[i%len(usernames)] + fmt.Sprintf("%d", i),
			Email:     usernames[i%len(usernames)] + fmt.Sprintf("%d", i) + "@example.com",
			Password:  "123",
			FirstName: usernames[i%len(usernames)] + fmt.Sprintf("%d", i),
			LastName:  usernames[i%len(usernames)] + fmt.Sprintf("%d", i),
			Phone:     "12321321",
			Address:   usernames[i%len(usernames)] + fmt.Sprintf("%d", i),
			Role: store.Role{
				Name: "user",
			},
		}
	}

	return users
}

func generateRoles(num int) []*store.Role {
	roles := make([]*store.Role, num)

	for i := 0; i < num; i++ {
		roles[i] = &store.Role{
			Name:        roleNames[i],
			Description: roleNames[i],
			Level:       i,
		}
	}

	return roles
}

func generateCategories(num int) []*store.Category {
	categories := make([]*store.Category, num)

	for i := 0; i < num; i++ {
		categories[i] = &store.Category{
			CategoryName: categoryNames[i%len(categoryNames)],
		}
	}

	return categories
}

func generateProducts(num int) []*store.Product {
	products := make([]*store.Product, num)

	for i := 0; i < num; i++ {
		products[i] = &store.Product{
			ID:          int64(i + 1),
			ProductName: productNames[i%len(productNames)],
			Description: productDescriptions[i%len(productDescriptions)],
			Price:       int64((i + 1) * 1000),
			Stock:       productStocks[i%len(productStocks)],
			Size:        productSize[i%len(productSize)],
			Color:       productColor[i%len(productColor)],
			CategoryID:  productCategoryIds[i%len(productCategoryIds)],
		}
	}

	return products
}
