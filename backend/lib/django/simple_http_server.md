# http.server - HTTP servers

One class, HTTPServer, is a socketserver.TCPServer subclass. It creates and listens at the HTTP socket, dispatching the requests to a handler. Code to create and run the server looks like this:

**class http.server.HTTPServer(server_address, RequestHandlerClass)**

This class builds on the TCPServer class by sorting the server address as instance variables named server_name and server_port. The server is accessible by the handler, typically through the handler's server instance variable.


**class http.server.ThreadingHTTPServer(server_address, RequestHandlerClass)**

This class is identical to HTTPServer but uses threads to handle requests by using the ThreadingMinxIn. This is useful to handle web browsers pre-opening sockets

```py
def run(server_class=HTTPServer, handler_class=BaseHTTPRequestHandler):
    server_address = ('', 8080)
    httpd = 
```

```py
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
```
