package cgibin

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("Cache-Control: no-cache")
	fmt.Println("Content-type: application/json")
	fmt.Println()

	jsonHello := map[string]string{
		"title":   "Hello, Go!",
		"heading": "Hello, Go!",
		"message": "This page was generated with the Go programming language.",
		"time":    time.Now().Format(time.RFC1123),
		"IP":      os.Getenv("REMOTE_ADDR"),
	}

	jsonData, err := json.Marshal(jsonHello)
	if err != nil {
		fmt.Println(`{"error": "Failed to generate JSON"}`)
		return
	}
	fmt.Println(string(jsonData))
}
