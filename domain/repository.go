package domain

import "context"

type BookRepo interface {
	CreateBook(ctx context.Context, b Book) (*Book, error)
	GetBook(ctx context.Context, name string) (*Book, error)
	ListBooks(ctx context.Context, limit, offset int64) ([]*Book, error)
	DeleteBook(ctx context.Context, id string) error
	UpdateBook(ctx context.Context, b Book) (*Book, error)
}
