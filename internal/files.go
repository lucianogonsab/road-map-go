package internal

import (
	"io"
	"os"
)

func OpenFile(fileName string) (string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	defer file.Close()

	return "Archivo abierto exitosamente", nil
}

func CreateFile(fileName string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString("Escribiendo en un archivo txt. \n")
	if err != nil {
		return err
	}

	return nil
}

func ReaderFile(fileName string) (string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}

	return string(data), nil
}
