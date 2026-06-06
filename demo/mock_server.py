#!/usr/bin/env python3
"""
Mock API server for apidiff demo.
Serves v1.json by default. Send a GET to /switch to flip to v2.json.
"""

from http.server import BaseHTTPRequestHandler, HTTPServer
import json, os

BASE = os.path.dirname(os.path.abspath(__file__))
state = {"version": "v1"}

class Handler(BaseHTTPRequestHandler):
    def do_GET(self):
        if self.path == "/switch":
            state["version"] = "v2"
            self._respond(200, {"switched": True, "now_serving": "v2.json"})
        elif self.path == "/user":
            fname = os.path.join(BASE, f"{state['version']}.json")
            with open(fname) as f:
                data = json.load(f)
            self._respond(200, data)
        else:
            self._respond(404, {"error": "not found"})

    def _respond(self, code, payload):
        body = json.dumps(payload, indent=2).encode()
        self.send_response(code)
        self.send_header("Content-Type", "application/json")
        self.send_header("Content-Length", str(len(body)))
        self.end_headers()
        self.wfile.write(body)

    def log_message(self, fmt, *args):
        version = state["version"]
        print(f"  [{version}] {self.path}")

if __name__ == "__main__":
    port = 8080
    print(f"Mock API running at http://localhost:{port}")
    print(f"  GET /user    → serves current version JSON")
    print(f"  GET /switch  → flips from v1 → v2")
    HTTPServer(("", port), Handler).serve_forever()
