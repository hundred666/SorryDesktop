package main

import (
	"testing"
	"fmt"
)

func TestParseAss(t *testing.T) {
	s, _ := ParseAss("E:\\develop\\goproject\\a.ass")
	fmt.Println(s)
}
