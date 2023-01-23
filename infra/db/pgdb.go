package db

import (
	"book-shop/domain"
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresRepo struct {
	db *pgxpool.Pool
}

var _ domain.BookRepo = (*PostgresRepo)(nil)

type PostgresSpecification interface {
	Create(pgxpool.Config) pgxpool.Config
}

func NewPostgresRepo(cfg *pgxpool.Config) (*PostgresRepo, error) {
	db, err := pgxpool.NewWithConfig(context.Background(), cfg)
	if err != nil {
		return nil, err
	}

	return &PostgresRepo{db: db}, nil
}

func (r *PostgresRepo) CreateBook(ctx context.Context, b domain.Book) (*domain.Book, error) {
	authorGetSql := `SELECT id FROM authors WHERE id = $1`
	authorSql := `INSERT INTO authors (id, name)VALUES($1, $2)`
	bookSql := `INSERT INTO books (id, name, author_id) VALUES($1, $2, $3)`

	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, err
	}

	var authorID string

	err = func() error {
		err = tx.QueryRow(ctx, authorGetSql, b.Author.Name).Scan(&authorID)
		if err != pgx.ErrNoRows {
			return err
		}

		if authorID == "" {
			authorID = uuid.NewString()
		}

		_, err = tx.Exec(ctx, authorSql, authorID, b.Author.Name)
		if err != nil {
			return err
		}

		_, err = tx.Exec(ctx, bookSql, b.ID, b.Name, authorID)
		if err != nil {
			return err
		}

		return nil
	}()

	if err != nil {
		tx.Rollback(ctx)
		return nil, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, err
	}

	return &b, nil
}
func (r *PostgresRepo) GetBook(ctx context.Context, name string) (*domain.Book, error) {
	panic("not implem")
}

func (r *PostgresRepo) ListBooks(ctx context.Context, limit, offset int64) ([]*domain.Book, error) {
	sql := `SELECT books.id, books.name, authors.name 
	FROM books LEFT JOIN authors ON books.author_id = authors.id limit $1 offset $2`

	rws, err := r.db.Query(ctx, sql, limit, offset)
	if err != nil {
		return nil, err
	}

	defer func() {
		if rws != nil {
			rws.Close()
		}
	}()

	var books []*domain.Book
	for rws.Next() {
		book := domain.Book{}
		author := domain.Author{}
		if err = rws.Scan(&book.ID, &book.Name, &author.Name); err != nil {
			return nil, err
		}

		book.Author = &author
		books = append(books, &book)
	}

	return books, nil
}

func (r *PostgresRepo) DeleteBook(ctx context.Context, id string) error { panic("not implem") }
func (r *PostgresRepo) UpdateBook(ctx context.Context, b domain.Book) (*domain.Book, error) {
	panic("not implem")
}
