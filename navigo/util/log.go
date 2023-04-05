package util

import (
	"log"
)

func AddRoute(path, method string) {
	log.Printf("(Path: %v, Method: %v)", path, method)
}

func LogRoute(path, method string) {
	log.Printf("ROUTE INFO (Path: %v, Method: %v)", path, method)
}
