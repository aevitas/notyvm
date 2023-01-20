package names

import (
	"errors"
	"fmt"
	"strings"
)

func GetEmailAddress(fn string, ln string) (string, error) {
	if len(fn) == 0 {
		return "", errors.New("first name is required")
	}

	if len(ln) == 0 {
		return "", errors.New("last name is required")
	}

	return strings.ToLower(fmt.Sprintf("%s.%s@isveiled.com", fn, ln)), nil
}
