package internal

import (
	"fmt"
	"io"
	"os"
)

type DivisionByZeroError struct {
	Message string
}

type UserNotFoundError struct {
	UserName string
}

type PasswordIncorrectError struct{}

func (e UserNotFoundError) Error() string {
	return fmt.Sprintf("user %s not found", e.UserName)
}

func (e PasswordIncorrectError) Error() string {
	return "incorrect password"
}

func (e DivisionByZeroError) Error() string {
	return e.Message
}

func ReadFile(fileName string) (string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}

	return string(content), nil
}

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, DivisionByZeroError{"No se puede dividir por cero"}
	}

	return a / b, nil
}

var Users = map[string]string{"Luciano": "contra", "Oscar": "pass"}

func Authentication(user, pass string) error {
	password, ok := Users[user]
	if !ok {
		return UserNotFoundError{"user"}
	}
	if password != pass {
		return PasswordIncorrectError{}
	}

	return nil
}

//dasdasda
