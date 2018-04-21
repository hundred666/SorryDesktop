package main

import (
	"testing"
	"fmt"
)

func TestGeneratemovie(t *testing.T) {
	words := make([]string, 0)
	for i := 0; i < 4; i++ {
		words = append(words, "abcd")
	}
	movieFile := "E:\\develop\\goproject\\template.mp4"
	assFile := "E:\\develop\\goproject\\template.ass"
	path, _ := GenerateMovie(movieFile, assFile, words)
	fmt.Println(path)
}
