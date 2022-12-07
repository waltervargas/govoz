package govoz

import (
	"encoding/binary"
	"log"
	"net/http"

	"github.com/gordonklaus/portaudio"
)

const (
	listenOn = ":8080"
)

func runHTTPServer() error {
	portaudio.Initialize()
	defer portaudio.Terminate()
	buffer := make([]float32, sampleRate*seconds)
	stream, err := portaudio.OpenDefaultStream(1, 0, sampleRate, len(buffer), func(in []float32) {
		for i := range buffer {
			buffer[i] = in[i]
		}
	})

	if err != nil {
		return err
	}

	if err = stream.Start(); err != nil {
		return err
	}
	defer stream.Close()

	http.HandleFunc("/audio", func(w http.ResponseWriter, r *http.Request) {
		flusher, ok := w.(http.Flusher)
		//TODO: (walter) to remove panic and handle error
		if !ok {
			panic("expected http.ResponseWriter to be an http.Flusher")
		}
		w.Header().Set("Connection", "Keep-Alive")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("Transfer-Encoding", "chunked")
		w.Header().Set("Content-Type", "audio/wave")
		for {
			log.Printf("connection received")
			binary.Write(w, binary.BigEndian, &buffer)
			flusher.Flush() // Trigger "chunked" encoding and send a chunk...
			return
		}
	})

	http.ListenAndServe(listenOn, nil)

	return nil
}
