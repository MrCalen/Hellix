package udp

import (
	"fmt"
	"net"
	"os"
)

func CreateServer(
	onMessage func([]byte),
) {
	conn, err := net.ListenPacket("udp", os.Getenv("REMOTE_SERVER"))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	for {
		msgBuffer := make([]byte, 1024)
		n, _, err := conn.ReadFrom(msgBuffer)
		if err != nil {
			fmt.Println("Err", err)
			continue
		}
		onMessage(msgBuffer[:n])
	}
}
