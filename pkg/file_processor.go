package pkg

import (
	"bytes"
	"errors"
	"io"
	"os"

	"github.com/hajimehoshi/go-mp3"
)

func ProcessFile(filename string) ([]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, errors.New("open file: file not found")
	}
	defer file.Close()

	// Decode MP3 to PCM
	decoder, err := mp3.NewDecoder(file)
	if err != nil {
		return nil, errors.New("mp3 decoder: can't decode mp3 file")
	}

	// Read decoded PCM data into a byte slice
	var buf bytes.Buffer
	_, err = io.Copy(&buf, decoder)
	if err != nil {
		return nil, errors.New("copy: can't copy decoder output")
	}

	datas := buf.Bytes()
	intData := make([]int, len(datas))
	for i, b := range datas {
		intData[i] = int(b)
	}

	return intData, nil
}
