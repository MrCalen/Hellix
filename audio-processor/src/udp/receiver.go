package udp

import (
	"fmt"
	"net"
	"os"
)

func CreateServer(
	onMessage func([]byte),
) {
	serverAdd, err := net.ResolveUDPAddr("udp", os.Getenv("REMOTE_SERVER"))
	if err != nil {
		panic(err)
	}
	conn, err := net.ListenUDP("udp", serverAdd)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	msgBuffer := make([]byte, 1024)
	for {
		n, _, err := conn.ReadFromUDP(msgBuffer)
		if err != nil {
			fmt.Println("Err", err)
			continue
		}
		onMessage(msgBuffer[:n])
	}
}
