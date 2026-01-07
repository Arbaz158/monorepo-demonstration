import json
import os
from http.server import BaseHTTPRequestHandler, HTTPServer


class RequestHandler(BaseHTTPRequestHandler):
    def _set_headers(self, status=200):
        self.send_response(status)
        self.send_header("Content-Type", "application/json")
        self.end_headers()

    def do_GET(self):
        if self.path == "/health":
            self._set_headers()
            self.wfile.write(json.dumps({"status": "ok", "service": "notification-service"}).encode())
        elif self.path == "/notify":
            self._set_headers()
            payload = {"delivered": True, "message": "Notification sent"}
            self.wfile.write(json.dumps(payload).encode())
        else:
            self._set_headers(404)
            self.wfile.write(json.dumps({"error": "not found"}).encode())

    def log_message(self, format, *args):  # noqa: A003
        # Keep stdout tidy for demo purposes.
        return


def run():
    port = int(os.getenv("PORT", "5001"))
    server = HTTPServer(("", port), RequestHandler)
    print(f"Starting notification-service on :{port}")
    server.serve_forever()


if __name__ == "__main__":
    run()

