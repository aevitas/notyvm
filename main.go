package main

func main() {
	s := &Server{}

	s.Init()
	s.Start(":8080")
}
