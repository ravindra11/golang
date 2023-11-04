package randomStr

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
)

func randomFormat() string {

	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! well met!",
	}
	return formats[rand.Intn(len(formats))]
}

func RandomString(name string) (string, error) {
	if strings.TrimSpace(name) == "" {
		return name, errors.New("empty name")
	}
	return fmt.Sprintf(randomFormat(), name), nil
}

/* multiple greetings to multiple people */

func MultipleGreetings(names []string) (map[string]string, error) {

	if names == nil {
		return nil, errors.New("Empty Names")
	}
	messages := make(map[string]string)

	for _, name := range names {

		message, err := RandomString(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}
