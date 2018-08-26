package main

import (
	"audio-processor/src/message"
	"audio-processor/src/udp"
	"audio-processor/src/ws"
)

func main() {

	websocketConnections := ws.OpenSockets{}

	go ws.CreateServer(&websocketConnections)

	udp.CreateServer(func(msg []byte) {
		message.OnMessage(msg, websocketConnections)
	})
}
