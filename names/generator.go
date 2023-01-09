package names

import (
	"log"
	"os"
	"strings"
)

func GenerateName(seed int) (string, string) {
	fn := getFirstNames()
	ln := getLastNames()

	return fn[seed%len(fn)], ln[seed%len(ln)]
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
