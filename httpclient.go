package govoz

import (
	"bytes"
	"encoding/binary"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gordonklaus/portaudio"
)

func runHTTPClient(url string) error {
	portaudio.Initialize()
	defer portaudio.Terminate()
	buffer := make([]float32, sampleRate*seconds)

	stream, err := portaudio.OpenDefaultStream(0, 1, sampleRate, len(buffer), func(out []float32) {
		resp, err := http.Get(url)
		if err != nil {
			// TODO: (walter) handle error
			panic(err)
		}
		body, _ := ioutil.ReadAll(resp.Body)
		responseReader := bytes.NewReader(body)
		binary.Read(responseReader, binary.BigEndian, &buffer)
		for i := range out {
			out[i] = buffer[i]
		}
	})
	if err != nil {
		return err
	}

	err = stream.Start()
	if err != nil {
		return err
	}

	// TODO: understand better this line
	time.Sleep(time.Second * 40)

	err = stream.Stop()
	if err != nil {
		return err
	}
	defer stream.Close()

	return nil
}
