package udp

import (
	"audio-collector/src/audio"
	"encoding/json"
	"fmt"
	"net"
	"os"
)

func ConnectToServer() net.Conn {
	url := os.Getenv("REMOTE_SERVER")

	conn, err := net.Dial("udp", url)
	if err != nil {
		panic(err)
	}

	return conn
}

func SendBuffer(buffer audio.Buffer, remoteConnection net.Conn) {
	json, err := json.Marshal(buffer)
	if err != nil {
		fmt.Println("Error encountered when send data")
		return
	}
	remoteConnection.Write(json)
}
