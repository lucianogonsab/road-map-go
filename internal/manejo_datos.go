package internal

import (
	"encoding/csv"
	"encoding/json"
	"os"
	"strconv"
)

type Person struct {
	Name  string
	Age   int
	Email string
}

type Human struct {
	ID   int
	Name string
	Age  int
}

func DecodeJson(p Person) ([]byte, error) {
	jsonData, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}

func UncodeJson(p []byte) (string, error) {
	var persona Person
	err := json.Unmarshal(p, &persona)
	if err != nil {
		return "", err
	}
	result := "El nombre es: " + persona.Name + " y tiene " + strconv.Itoa(persona.Age)
	return result, nil
}

func ReadArchive(name string) (*os.File, error) {
	archivo, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	return archivo, nil
}

func ConvertToPersonStruct(arch *os.File) ([]Human, error) {
	defer arch.Close()

	lector := csv.NewReader(arch)
	registros, err := lector.ReadAll()
	if err != nil {
		return []Human{}, err
	}

	var person = []Human{}
	for _, reg := range registros {
		id, _ := strconv.Atoi(reg[0])
		age, _ := strconv.Atoi(reg[2])
		person = append(person, Human{ID: id, Name: reg[1], Age: age})
	}

	return person, nil
}

func FilterHumanByAge(people []Human, filtro int) []Human {
	peopleFilter := []Human{}
	for _, p := range people {
		if p.Age > filtro {
			peopleFilter = append(peopleFilter, p)
		}
	}

	return peopleFilter
}
