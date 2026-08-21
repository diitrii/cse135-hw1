package cgibin

import (
	"fmt"
	"net/url"
	"os"
	"strings"
	"time"
)

func main() {
	os.MkdirAll("/tmp/go-sessions", 0700)
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
	if sessionID == "" {
		sessionID = fmt.Sprintf("%d", time.Now().UnixNano())
	}

	sessionFile := "/tmp/go-sessions/" + sessionID

	username := ""
	data, err := os.ReadFile(sessionFile)
	if err == nil {
		username = string(data)
	}

	if username == "" {
		contentLength := os.Getenv("CONTENT_LENGTH")
		if contentLength != "" {
			var length int
			fmt.Sscanf(contentLength, "%d", &length)
			data := make([]byte, length)
			os.Stdin.Read(data)
			values, err := url.ParseQuery(string(data))
			if err == nil {
				username = values.Get("username")
			}
		}
	}
	if username != "" {
		os.WriteFile(sessionFile, []byte(username), 0600)
	}

	fmt.Println("Content-Type: text/html")
	fmt.Printf("Set-Cookie: GOSESSIONID=%s; Path=/\r\n", sessionID)
	fmt.Println()

	fmt.Println("<html>")
	fmt.Println("<head>")
	fmt.Println("<title>Go Sessions</title>")
	fmt.Println("</head>")
	fmt.Println("<body>")
	fmt.Println("<h1>Go Sessions Page 1</h1>")

	if username != "" {
		fmt.Printf("<p><b>Name:</b> %s</p>\n", username)
	} else {
		fmt.Println("<p><b>Name:</b> You do not have a name set</p>")
	}

	fmt.Println("<br/><br/>")
	fmt.Println("<a href='/cgi-bin/go-sessions-2'>Session Page 2</a><br/>")
	fmt.Println("<a href='/go-cgiform.html'>Go CGI Form</a><br />")
	fmt.Println("<form style='margin-top:30px' action='/cgi-bin/go-destroy-session' method='get'>")
	fmt.Println("<button type='submit'>Destroy Session</button>")
	fmt.Println("</form>")
	fmt.Println("</body>")
	fmt.Println("</html>")
}
