package files

import (
	_ "embed"
)

//go:embed mocks/files.json
var filesJSON []byte
