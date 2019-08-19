package audio

import (
	"github.com/go-audio/audio"
	"github.com/go-audio/generator"
	"github.com/gordonklaus/portaudio"
	"log"
)

func Dot() {

}

func Dash() {

}

func makeNoise() {
	bufferSize := 512
	frequency := 440.0

	buf := &audio.FloatBuffer{
		Data:   make([]float64, bufferSize),
		Format: audio.FormatMono22500,
	}

	osc := generator.NewOsc(generator.WaveSine, frequency, buf.Format.SampleRate)
	osc.Amplitude = 1

	portaudio.Initialize()
	defer portaudio.Terminate()

	out := make([]float32, bufferSize)

	stream, err := portaudio.OpenDefaultStream(0, 1, 44100, len(out), &out)
	if err != nil {
		log.Fatal(err)
	}

	defer stream.Close()

	if err := stream.Start(); err != nil {
		log.Fatal(err)
	}

	defer stream.Stop()

	for {
		if err := osc.Fill(buf); err != nil {
			log.Fatal(err)
		}

		f64ToF32Copy(out, buf.Data)
	}

	if err := stream.Write(); err != nil {
		log.Fatal(err)
	}
}

func f64ToF32Copy(dst []float32, src []float64) {
	for i := range src {
		dst[i] = float32(src[i])
	}
}
