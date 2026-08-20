package cgibin

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Cache-Control: no-cache")
	fmt.Println("Content-type: text/html")
	fmt.Println()
	fmt.Println(`<!DOCTYPE html><html><head><title>General Request Echo</title></head><body><h1 align="center">General Request Echo</h1><hr>`)
	fmt.Printf("<p><b>HTTP Protocol:</b> %s</p>\n", os.Getenv("SERVER_PROTOCOL"))
	fmt.Printf("<p><b>HTTP Method:</b> %s</p>\n", os.Getenv("REQUEST_METHOD"))
	fmt.Printf("<p><b>Query String:</b> %s</p>\n", os.Getenv("QUERY_STRING"))

	contentLength := os.Getenv("CONTENT_LENGTH")

	var formData []byte

	if contentLength != "" {
		var length int
		fmt.Sscanf(contentLength, "%d", &length)

		formData = make([]byte, length)
		_, err := io.ReadFull(os.Stdin, formData)

		if err != nil {
			fmt.Printf("<p><b>Error reading body:</b> %s</p>\n", err)
		}
	}

	fmt.Printf("<p><b>Message Body:</b> %s</p>\n", string(formData))
	fmt.Println("</body></html>")
}
