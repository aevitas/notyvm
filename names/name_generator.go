package names

import (
	"log"
	"os"
	"strings"
)

type Name struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func RandomFirstName(seed int) string {
	n := getFirstNames()

	return n[seed%len(n)]
}

func RandomLastName(seed int) string {
	n := getLastNames()

	return n[seed%len(n)]
}

func RandomName(seed int) *Name {
	fn := getFirstNames()
	ln := getLastNames()

	return &Name{FirstName: fn[seed%len(fn)], LastName: ln[seed%len(ln)]}
}

func getFirstNames() []string {
	b, err := os.ReadFile("first.txt")

	if err != nil {
		log.Fatal(err)
	}

	fn := strings.Split(string(b), "\r\n")

	return fn
}

func getLastNames() []string {
	b, err := os.ReadFile("last.txt")

	if err != nil {
		log.Fatal(err)
	}

	ln := strings.Split(string(b), "\r\n")

	return ln
}
