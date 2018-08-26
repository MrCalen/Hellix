package ws

import (
	"net/http"

	"github.com/gorilla/websocket"
)

type OpenSockets = []*websocket.Conn

func CreateServer(openConnections *OpenSockets) {
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		handleWebsocketConnection(w, r, openConnections)
	})
	err := http.ListenAndServe(":1338", nil)
	if err != nil {
		panic(err)
	}
}

func handleWebsocketConnection(w http.ResponseWriter, r *http.Request, openConnections *OpenSockets) {
	conn, err := websocket.Upgrade(w, r, w.Header(), 1024, 1024)
	if err != nil {
		http.Error(w, "Error", http.StatusBadRequest)
		return
	}

	*openConnections = append(*openConnections, conn)
}
