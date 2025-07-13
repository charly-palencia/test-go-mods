package shares

import (
	"log"
)

func LogRequest(service, action string) {
	log.Printf("[service: %s] Action: %s", service, action)
}
