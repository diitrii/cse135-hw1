#!usr/bin/env python3

import json
import time
import os

print("Cache-Control: no-cache\n")
print("Content-Type: application/json\n\n")

date = time.localTime()
address = os.environ.get("REMOTE_ADDR")
data = {
    "title": "Hello, HyperTrees!",
    "heading": "Hello, HyperTrees!",
    "message": "This message was generated with the Python programming language",
    "time": date,
    "IP": address
}

json_string = json.dumps(data)
print(json_string)

pretty_json_string = json.dumps(data, indent=2)
print(pretty_json_string)
