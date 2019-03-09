package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/seregayoga/factorial/pkg/handler"
)

func main() {
	http.HandleFunc("/calculate", handler.CalculateHandler)

	http.ListenAndServe(getAddress(), nil)
}

func getAddress() string {
	return fmt.Sprintf("%s:%s", getEnv("FACTORIAL_HOST", ""), getEnv("FACTORIAL_PORT", "5000"))
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
