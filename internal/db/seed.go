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

func Seed(store store.Storage, db *sql.DB) {
	ctx := context.Background()

	// roles := generateRoles(2)

	users := generateUsers(5)

	tx, _ := db.BeginTx(ctx, nil)

	// for _, role := range roles {
	// 	err := store.Roles.Create(ctx, tx, role)
	// 	if err != nil {
	// 		_ = tx.Rollback()
	// 		log.Println("Error creating role:", err)
	// 		return
	// 	}
	// }

	result, err := store.Users.GetAll(ctx)
	if err != nil {
		fmt.Println(err)
	}

	for _, item := range result {
		fmt.Println(item.Username)
	}

	for _, user := range users {
		err := store.Users.Create(ctx, tx, user)
		if err != nil {
			_ = tx.Rollback()
			log.Println("Error creating user:", err)
			return
		}

	}

	tx.Commit()

	log.Println("Seeding complete")
}

func generateUsers(num int) []*store.User {
	users := make([]*store.User, num)

	for i := 0; i < num; i++ {
		users[i] = &store.User{
			Username:  usernames[i%len(usernames)] + fmt.Sprintf("%d", i),
			Email:     usernames[i%len(usernames)] + fmt.Sprintf("%d", i) + "@example.com",
			AvatarUrl: "test link",
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
