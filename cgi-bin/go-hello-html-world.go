package main

import (
    "fmt"
    "os"
    "time"
)

func main(){
    fmt.Println("Cache-Control: no-cache")
    fmt.Println("Content-type: text/html")
    fmt.Println()
    fmt.Println("<!DOCTYPE html>")
    fmt.Println("<html>")
    fmt.Println("<head>")
    fmt.Println("<title>Hello CGI World</title>")
    fmt.Println("</head>")
    fmt.Println("<body>")
    fmt.Println("<h1 align=\"center\">Hello Go World</h1><hr/>")
    fmt.Println("<p>Hello World</p>")
    fmt.Println("<p>This page was generated with the Go programming language.</p>")

    fmt.Printf("<p>This program was generated at: %s.</p>\n", time.Now().Format(time.RFC1123))
    fmt.Printf("<p>Your current IP Address is: %s</p>\n", os.Getenv("REMOTE_ADDR"))
    fmt.Println("</body>")
    fmt.Println("</html>")
}