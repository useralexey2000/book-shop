package db

import (
	"book-shop/domain"
	"context"
	"encoding/json"
	"fmt"
)

type MockDB struct{}

func (r *MockDB) CreateBook(ctx context.Context, b domain.Book) (*domain.Book, error) {
	panic("not impl")
}
func (repo *MockDB) GetBook(ctx context.Context, name string) (*domain.Book, error) {
	panic("not impl")
}
func (r *MockDB) ListBooks(ctx context.Context, limit, offset int64) ([]*domain.Book, error) {
	js := `[{"ID":"928418ca-7d1e-4bce-a961-270468a0e9dc","Name":"book1","Author":{"Name":"book1"}},{"ID":"d2a80407-63d1-4882-a036-f84ab657dd22","Name":"book2","Author":{"Name":"book2"}}]`

	var bl []*domain.Book
	err := json.Unmarshal([]byte(js), &bl)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return bl, nil
}
func (r *MockDB) DeleteBook(ctx context.Context, id string) error { panic("not impl") }
func (r *MockDB) UpdateBook(ctx context.Context, b domain.Book) (*domain.Book, error) {
	panic("not impl")
}
