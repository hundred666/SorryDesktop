package main

import (
	"github.com/asticode/go-astilectron"
	"encoding/json"
	"github.com/asticode/go-astilectron-bootstrap"
	"strings"
)

// handleMessages handles messages
func handleMessages(_ *astilectron.Window, m bootstrap.MessageIn) (payload interface{}, err error) {
	switch m.Name {
	case "load":
		// Unmarshal payload
		var path []string
		if len(m.Payload) > 0 {
			// Unmarshal payload
			if err = json.Unmarshal(m.Payload, &path); err != nil {
				payload = err.Error()
				return
			}
		}
		// Explore
		if payload, err = loadAss(path[0]); err != nil {
			payload = err.Error()
			return
		}
	case "generate":
		var result Result
		if len(m.Payload) > 0 {
			// Unmarshal payload
			if err = json.Unmarshal(m.Payload, &result); err != nil {
				payload = err.Error()
				return
			}
		}
		if payload, err = generate(result); err != nil {
			payload = err.Error()
			return
		}
	}
	return
}

type Subtitle struct {
	Words []string `json:"words"`
	Len   int      `json:"len"`
}

type Result struct {
	Words  []string `json:"words"`
	Film   string   `json:"film"`
	Output string   `json:"output"`
}

func loadAss(path string) (s Subtitle, err error) {
	if len(path) == 0 {
		return
	}

	subtitle := strings.Split(path, ".")[0] + ".ass"
	//parse ass file to find length of words
	s, err = ParseAss(subtitle)
	return
}

func generate(result Result) (path string, err error) {
	words := result.Words
	filmFile :=string([]rune(result.Film)[7:])
	assFile := strings.Split(filmFile, ".")[0] + ".ass"
	output, err := GenerateMovie(filmFile, assFile, words)
	return output, err
}
