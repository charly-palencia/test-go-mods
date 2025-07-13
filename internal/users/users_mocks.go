package users

import (
	_ "embed"
)

//go:embed mocks/users.json
var usersJSON []byte
