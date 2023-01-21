package infra

import (
	"book-shop/domain"
	"book-shop/mapper"
	"book-shop/proto/pb"
	"context"

	"github.com/google/uuid"
)

type LibraryService struct {
	bookRepo domain.BookRepo
	pb.UnsafeBookServiceServer
}

func NewLibraryService(repo domain.BookRepo) *LibraryService {
	return &LibraryService{
		bookRepo: repo,
	}
}

func (l *LibraryService) CreateBook(
	ctx context.Context, req *pb.CreateBookRequest) (*pb.CreateBookResponse, error) {

	book := domain.Book{
		ID:   uuid.NewString(),
		Name: req.Name,
		Author: &domain.Author{
			Name: req.AuthorName,
		},
	}

	b, err := l.bookRepo.CreateBook(ctx, book)
	if err != nil {
		return &pb.CreateBookResponse{}, err
	}

	protoBook := mapper.BookToProto(b)
	return &pb.CreateBookResponse{Book: protoBook}, nil
}

func (l *LibraryService) ListBooks(
	ctx context.Context, req *pb.ListBooksRequest) (*pb.ListBooksResponse, error) {

	books, err := l.bookRepo.ListBooks(ctx, req.Limit, req.Offset)
	if err != nil {
		return &pb.ListBooksResponse{}, err
	}

	protoBooks := mapper.BookListToProto(books)
	res := &pb.ListBooksResponse{Books: protoBooks}
	return res, nil
}
