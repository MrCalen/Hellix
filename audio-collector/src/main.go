package main

import (
	"audio-collector/src/audio"
	"audio-collector/src/udp"
)

func main() {
	remoteConnection := udp.ConnectToServer()
	audio.Collect(func(buffer audio.Buffer) {
		udp.SendBuffer(buffer, remoteConnection)
	})
}
