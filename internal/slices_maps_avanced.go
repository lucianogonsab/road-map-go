package internal

import "errors"

type TypeInteger struct {
	Number int
}

type TypeString struct {
	Word string
}

type Book struct {
	Title  string
	Author string
	Year   int64
}

func JoinSlice(a []TypeInteger, b []TypeString) []interface{} {
	sliceCombined := make([]interface{}, len(a)+len(b))

	for i, integer := range a {
		sliceCombined[i] = integer
	}

	for i, word := range b {
		sliceCombined[i+len(a)] = word
	}

	return sliceCombined
}

func AddBook(library map[string]Book, book Book) error {
	_, exists := library[book.Title]
	if exists {
		return errors.New("This book already exists")
	}
	library[book.Title] = book
	return nil
}

func GetBook(library map[string]Book, title string) (error, Book) {
	value, exists := library[title]
	if !exists {
		return errors.New("This book does not exist"), value
	}
	return nil, value
}

func EditBook(library map[string]Book, title string, book Book) (error, Book) {
	_, exists := library[title]
	if !exists {
		return errors.New("This book does not exist"), book
	}
	library[title] = book
	return nil, book
}

func DeleteBook(library map[string]Book, title string) error {
	_, exists := library[title]
	if !exists {
		return errors.New("Not exists any book with that title")
	}
	delete(library, title)
	return nil
}
