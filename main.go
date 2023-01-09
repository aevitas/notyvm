package main

import "aevitas.dev/veiled/api"

func main() {
	s := &api.Server{}

	s.Init()
	s.Start(":8080")
}
