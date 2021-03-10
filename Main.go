package main

import (
	"gopkg.in/macaron.v1"
)

func main() {
	m := macaron.New()
	m.Use(macaron.Static("public"))
	m.Run()
}