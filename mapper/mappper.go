package mapper

import (
	"book-shop/domain"
	"book-shop/proto/pb"
)

func BookToProto(b *domain.Book) *pb.Book {
	return &pb.Book{
		ID:   b.ID,
		Name: b.Name,
		Author: &pb.Author{
			Name: b.Author.Name,
		},
	}
}

func ProtoToBook(b *pb.Book) *domain.Book {
	return &domain.Book{
		ID:   b.ID,
		Name: b.Name,
		Author: &domain.Author{
			Name: b.Name,
		},
	}
}

func BookListToProto(bs []*domain.Book) []*pb.Book {
	books := make([]*pb.Book, 0, len(bs))

	for _, v := range bs {
		b := BookToProto(v)
		books = append(books, b)
	}

	return books
}

func ProtoToListBook(bs []*pb.Book) []*domain.Book {
	books := make([]*domain.Book, 0, len(bs))

	for _, v := range bs {
		b := ProtoToBook(v)
		books = append(books, b)
	}

	return books
}
