package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"text/template"
)

/*
Solution hiện tại là dùng poller để check QR code status.
+ Lợi ích: dễ dàng implement
+ Nhược điểm: call nhiều

WebSocket cho phép two-way communication giữa client và web server.
+ Nhược điểm: khó implement hơn poller, phải maintain connection logic

+ Tạo api handler /api/v1/authn/qrcode/ws-complete
+ Mở ws trên path này với connection id là qrcode_id
+ Lib: https://github.com/gorilla/websocket
+ Xử lý các logic như api qrcode/complete
+ Đóng ws connection khi QR code status success, client close tab.
*/

var addr = flag.String("addr", ":8080", "http service address")

func main() {
	flag.Parse()
	log.SetFlags(0)
	hub := NewHub()
	go hub.run()
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})
	http.HandleFunc("/event", func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query().Get("msg")
		qrcodeID := r.URL.Query().Get("qrcode_id")
		hub.event <- &Event{
			ClientID: qrcodeID,
			Message:  []byte(q),
		}
		w.Write([]byte("publish event success"))
	})
	http.HandleFunc("/", home)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	qrcodeID := r.URL.Query().Get("qrcode_id")
	homeTemplate.Execute(w, fmt.Sprintf("ws://"+r.Host+"/ws?qrcode_id=%s", qrcodeID))
}

var homeTemplate = template.Must(template.New("").Parse(`
<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<script>  
window.addEventListener("load", function(evt) {

    var output = document.getElementById("output");
    var input = document.getElementById("input");
    var ws;

    var print = function(message) {
        var d = document.createElement("div");
        d.textContent = message;
        output.appendChild(d);
        output.scroll(0, output.scrollHeight);
    };

    document.getElementById("open").onclick = function(evt) {
        if (ws) {
            return false;
        }
        ws = new WebSocket("{{.}}");
        ws.onopen = function(evt) {
            print("OPEN");
        }
        ws.onclose = function(evt) {
            print("CLOSE");
            ws = null;
        }
        ws.onmessage = function(evt) {
            print("RESPONSE: " + evt.data);
        }
        ws.onerror = function(evt) {
            print("ERROR: " + evt.data);
        }
        return false;
    };

    document.getElementById("send").onclick = function(evt) {
        if (!ws) {
            return false;
        }
        print("SEND: " + input.value);
        ws.send(input.value);
        return false;
    };

    document.getElementById("close").onclick = function(evt) {
        if (!ws) {
            return false;
        }
        ws.close();
        return false;
    };

});
</script>
</head>
<body>
<table>
<tr><td valign="top" width="50%">
<p>Click "Open" to create a connection to the server, 
"Send" to send a message to the server and "Close" to close the connection. 
You can change the message and send multiple times.
<p>
<form>
<button id="open">Open</button>
<button id="close">Close</button>
<p><input id="input" type="text" value="Hello world!">
<button id="send">Send</button>
</form>
</td><td valign="top" width="50%">
<div id="output" style="max-height: 70vh;overflow-y: scroll;"></div>
</td></tr></table>
</body>
</html>
`))
