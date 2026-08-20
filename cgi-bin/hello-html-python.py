import time
import os
print("Cache-Control: no-cache\n")
print("Content-Type: text/html")

print("<!DOCTYPE html>")
print("<html>")
print("<head>")
print("<title>Hello HyperTrees World</title>")
print("</head>")
print("<body>")

print("<h1 align=center>Hello HTML World</h1><hr/>")
print("<p>HyperTrees</p>")
print("<p>This page was generated with the Python programming language</p>")

date = time.localTime();
print(f"<p>This program was generated at: {time}</p>")

address = os.environ.get("REMOTE_ADDR")
print(f"<p>Your current IP Address is: {address}</p>")

print("</body>")
print("</html>")

