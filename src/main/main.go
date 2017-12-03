package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
)

const port string = "8081"

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func ping(w http.ResponseWriter, r *http.Request) {
	Info.Println("Ping endpoint called")
	fmt.Fprint(w, "Service is Alive!")
}

func info(w http.ResponseWriter, r *http.Request) {
	Info.Println("Info endpoint called")
	fmt.Fprintf(w, "This is the Go server. Status: RUNNING. Port:"+port)
}

func testWs(w http.ResponseWriter, r *http.Request) {
	Info.Println("Index endpoint called")
	http.ServeFile(w, r, "html/index.html")
}

func main() {
	Init(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)
	Info.Println("<< Playlist Events Service - Started >>")

	// PATH HANDLES
	http.HandleFunc("/ping", ping)
	http.HandleFunc("/info", info)
	http.HandleFunc("/ws", ws)

	// Client WS for testing purposes
	http.HandleFunc("/test-ws", testWs)

	http.ListenAndServe(":"+port, nil)
}
