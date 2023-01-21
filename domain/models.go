package domain

type Book struct {
	ID     string
	Name   string
	Author *Author
}

func NewBook(id, name string, author *Author) *Book {
	return &Book{
		ID:     id,
		Name:   name,
		Author: author,
	}
}

type Author struct {
	Name string
}

func NewAuthor(name string) *Author {
	return &Author{
		Name: name,
	}
}
