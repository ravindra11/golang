package greetings

// declare package greetings to collect related functions

import (
	"errors"
	"fmt"
	"strings"
)

func Hello(name string) (string, error) {

	if strings.TrimSpace(name) == "" {
		return "", errors.New("empty name")
	}
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}
