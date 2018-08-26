package audio

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/gordonklaus/portaudio"
)

type Buffer = []float32

func Collect(afterCollect func(Buffer)) {
	portaudio.Initialize()
	defer portaudio.Terminate()
	in := make(Buffer, 64)
	stream, err := portaudio.OpenDefaultStream(1, 0, 44100, len(in), in)
	errCheck(err)
	defer stream.Close()

	ticker := []string{
		"-",
		"\\",
		"/",
		"|",
	}
	rand.Seed(time.Now().UnixNano())

	errCheck(stream.Start())
	for {
		fmt.Print("\r", ticker[rand.Intn(len(ticker)-1)])
		errCheck(stream.Read())
		buffer := in
		afterCollect(buffer)
	}
}

func errCheck(err error) {
	if err != nil {
		panic(err)
	}
}
