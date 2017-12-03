package main

import "net/http"

func ws(w http.ResponseWriter, r *http.Request) {
	Info.Println("New stablished connection")
	conn, _ := upgrader.Upgrade(w, r, nil) // error ignored for sake of simplicity

	for {
		// Read message from browser
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			return
		}

		// Print the message to the console
		Info.Println("Message received: " + string(msg))
		//HERE we process the json

		// Write message back to browser
		if err = conn.WriteMessage(msgType, msg); err != nil {
			return
		}
	}
}
