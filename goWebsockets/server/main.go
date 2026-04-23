// from tutorial https://www.youtube.com/watch?v=3C2RQgLCUJE

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/coder/websocket"
)

func echoHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := websocket.Accept(w, r, &websocket.AcceptOptions{
		InsecureSkipVerify: true, // Disable origin verification
	})
	if err != nil {
		log.Println("Failed to accept connection:", err)
		return
	}
	defer conn.Close(websocket.StatusNormalClosure, "")

	for {
		msgType, data, err := conn.Read(r.Context())
		if websocket.CloseStatus(err) != -1 {
			log.Println("Connection closed:", err)
			return
		}
		if err != nil {
			log.Println("Read error:", err)
			return
		}

		log.Printf("Received: %s\n", data)

		err = conn.Write(r.Context(), msgType, data)
		if err != nil {
			log.Println("Write error:", err)
			return
		}
	}
}

func main() {
	http.HandleFunc("/ws", echoHandler)
	fmt.Println("Echo server started at ws://localhost:8080/ws")
	log.Fatal(http.ListenAndServe(":8080", nil))
}