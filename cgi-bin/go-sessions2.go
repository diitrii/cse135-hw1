package main

import (
	"fmt"
	"net/url"
	"os"
	"strings"
)

func main() {
	cookieHeader := os.Getenv("HTTP_COOKIE")
	sessionID := ""
	cookies := strings.Split(cookieHeader, ";")

	for _, cookie := range cookies {
		parts := strings.SplitN(strings.TrimSpace(cookie), "=", 2)

		if len(parts) == 2 && parts[0] == "GOSESSIONID" {
			sessionID, _ = url.QueryUnescape(parts[1])
			break
		}
	}

	username := ""
	if sessionID != "" {
		sessionFile := "/tmp/go-sessions/" + sessionID
		data, err := os.ReadFile(sessionFile)
		if err == nil {
			username = string(data)
		}
	}

	fmt.Println("Content-Type: text/html")
	fmt.Println()
	fmt.Println("<html>")
	fmt.Println("<head>")
	fmt.Println("<title>Go Sessions</title>")
	fmt.Println("</head>")
	fmt.Println("<body>")
	fmt.Println("<h1>Go Sessions Page 2</h1>")

	if username != "" {
		fmt.Println("<p><b>Name:</b>", username, "</p>")
	} else {
		fmt.Println("<p><b>Name:</b> You do not have a name set</p>")
	}

	fmt.Println("<br/><br/>")
	fmt.Println("<a href='/cgi-bin/go-sessions-1'>Session Page 1</a><br/>")
	fmt.Println("<a href='/go-cgiform.html'>Go CGI Form</a><br />")
	fmt.Println("<form style='margin-top:30px' action='/cgi-bin/go-destroy-session' method='get'>")
	fmt.Println("<button type='submit'>Destroy Session</button>")
	fmt.Println("</form>")
	fmt.Println("</body>")
	fmt.Println("</html>")
}
