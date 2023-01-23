package gw

import (
	"book-shop/mapper"
	"book-shop/proto/pb"
	"encoding/json"
	"fmt"
	"net/http"
)

type Gateway struct {
	client pb.BookServiceClient
}

func New(c pb.BookServiceClient) *Gateway {
	return &Gateway{
		client: c,
	}
}

func (g *Gateway) CreateBook() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		type Req struct {
			Name       string
			AuthorName string
		}

		var req Req
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			return
		}

		res, err := g.client.CreateBook(r.Context(), &pb.CreateBookRequest{
			Name:       req.Name,
			AuthorName: req.AuthorName,
		})

		if err != nil {
			JSONResponse(w, err, http.StatusInternalServerError)
			return
		}

		JSONResponse(w, res, http.StatusOK)
	}
}

func (g *Gateway) ListBooks() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		type Req struct {
			Limit  int64 `json:"limit"`
			Offset int64 `json:"offset"`
		}

		var req Req
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			JSONResponse(w, err, http.StatusInternalServerError)
			return
		}

		res, err := g.client.ListBooks(r.Context(), &pb.ListBooksRequest{
			Limit:  req.Limit,
			Offset: req.Offset,
		})

		if err != nil {
			JSONResponse(w, err, http.StatusInternalServerError)
			return
		}

		books := mapper.ProtoToListBook(res.Books)
		JSONResponse(w, books, http.StatusOK)
	}
}

func JSONResponse(w http.ResponseWriter, val interface{}, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(val); err != nil {
		fmt.Println(err)
	}
}
