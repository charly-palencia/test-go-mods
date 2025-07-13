package files

import (
	"example.com/monorepo/internal/shares"
)

func ListFiles() []string {
	shares.LogRequest("files", "list")
	return shares.ReadFromBytes(filesJSON)
}
