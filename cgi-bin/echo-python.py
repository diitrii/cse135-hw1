#!usr/bin/env python3

import os

print("Cache-Control: no-cache\n")
print("Content-type: text/html\n\n")

print("<!DOCTYPE html>")
print("<html><head><title>General Request Echo</title>")
print('</head><body><h1 align="center">General Request Echo</h1>')
print("<hr>")

server_protocol = os.environ.get('SERVER_PROTOCOL', '')
request_method = os.environ.get('REQUEST_METHOD', '')
query_string = os.environ.get('QUERY_STRING', '')
print(f"Server Protocol: {server_protocol}")
print(f"Request Method: {request_method}")
print(f"Query String: {query_string}")

content_length = int(os.environ.get('CONTENT_LENGTH', 0) or 0)
form_data = sys.stdin.read(content_length)
print("<info-box>")
print(f"<p><b>Message Body:</b> {form_data}</p>")
print("</info-box>")
