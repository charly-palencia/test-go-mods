package shares

import (
	"encoding/json"
	"log"
)

func ReadFromBytes(data []byte) []string {
	var list []string
	if err := json.Unmarshal(data, &list); err != nil {
		log.Printf("error unmarshalling: %v", err)
		return []string{}
	}
	return list
}
