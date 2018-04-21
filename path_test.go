package main

import (
	"testing"
	"fmt"
)

func TestGetCurrentDirectory(t *testing.T) {
	p:=GetDirectory(".")
	fmt.Println(p)
}
