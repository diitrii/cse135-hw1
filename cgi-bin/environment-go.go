package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	fmt.Println("Cache-Control: no-cache")
	fmt.Println("Content-type: text/html")
	fmt.Println()
	fmt.Println("<!DOCTYPE html>")
	fmt.Println("<html>")
	fmt.Println("<head>")
	fmt.Println("<title>Environment Variables</title>")
	fmt.Println("</head>")
	fmt.Println("<body>")
	fmt.Println("<h1 align=\"center\">Environment Variables</h1><hr/>")
	fmt.Println("<p>")

	variables := os.Environ()
	sort.Strings(variables)
	for _, variable := range variables {
		vars := strings.SplitN(variable, "=", 2)
		fmt.Printf("<b>%s:</b> %s</p>\n", vars[0], vars[1])
	}

	fmt.Println("</p>")
	fmt.Println("</body>")
	fmt.Println("</html>")
}
