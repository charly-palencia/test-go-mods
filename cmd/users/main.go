package main

import (
	"fmt"
	"example.com/monorepo/internal/users"
)

func main() {
	usersList := users.ListUsers()
	fmt.Printf("✅ Users service: %d users found\n", len(usersList))
	for _, u := range usersList {
		fmt.Println("-", u)
	}
}
