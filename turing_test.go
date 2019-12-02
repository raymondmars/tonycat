package main

import (
	"log"
	"testing"
)

func TestChat(t *testing.T) {
	txt := (new(TuringBot)).Chat("你是谁")
	log.Println(txt)
}
