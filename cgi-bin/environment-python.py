#!usr/bin/env python3

import os

print("Cache-Control: no-cache\n")
print("Content-type: text/html\n\n")

print("<!DOCTYPE html>")
print("<html><head><title>Environment Variables</title>")
print('</head><body><h1 align="center">Environment Variables</h1>')
print("<hr>")

for param in os.environ.keys():
    print("<b>{}</b>: {}<br>".format(param, os.environ[param]))

print("</body></html")
