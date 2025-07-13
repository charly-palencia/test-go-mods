package users

import (
	"example.com/monorepo/internal/shares"
)

func ListUsers() []string {
	shares.LogRequest("users", "list")
	return shares.ReadFromBytes(usersJSON)
}
