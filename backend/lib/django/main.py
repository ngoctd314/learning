from http.server import BaseHTTPRequestHandler, HTTPServer
import json

host = "localhost"
port = 8080


class Server(BaseHTTPRequestHandler):
    def do_GET(self):
        self.send_response(200)
        self.send_header("Content-Type", "application/json")
        self.end_headers()
        data = json.dumps({"data": "pong"}).encode("utf-8")

        self.wfile.write(bytes(data))


if __name__ == "__main__":
    server = HTTPServer((host, port), Server)
    print("Server started http://%s:%s" % (host, port))

    try:
        server.serve_forever()
    except KeyboardInterrupt:
        pass

    server.server_close()
    print("Server stopped.")
