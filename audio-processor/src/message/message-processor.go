package message

import (
	"audio-collector/src/audio"
	"audio-processor/src/ws"
	"encoding/json"
	"fmt"
)

type Message = audio.Buffer

func OnMessage(jsonMsg []byte, sockets ws.OpenSockets) {
	var message Message

	err := json.Unmarshal(jsonMsg, &message)
	if err != nil {
		fmt.Println("JSON Error", err)
		return
	}

	for _, socket := range sockets {
		socket.WriteJSON(message)
	}
	fmt.Println(message)
}
