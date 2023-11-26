package childGreetings

import (
	"errors"
	"fmt"
	"math/rand"
	// "example.local/greetings"
)

var _ = fmt.Append // todo
var _ = errors.As //todo

func Hello2(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
	// message, err := greetings.Hello(name)
	// return message, err

}

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)
	for _, name := range names {
		message, err := Hello2(name)
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
