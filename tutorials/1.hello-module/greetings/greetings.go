package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	// "example.local/greetings/childGreetings"
)

var _ = errors.As // todo
var _ = fmt.Append // todo

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
	// message, err := childGreetings.Hello2(name)
	// return message, err

}

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}

func randomFormat() string {
	formats := [4]string{
		"Hi, %v Welcome",
		"Great to see you, %v!",
		"Hi, %v! Well Well",
	}

	return formats[rand.Intn(len(formats)-1)]
}
