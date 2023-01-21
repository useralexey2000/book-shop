package main

import (
	"book-shop/infra"
	"book-shop/mapper"
	"book-shop/proto/pb"
	"encoding/json"
	"fmt"
	"net/http"
)

type Gateway struct {
	service *infra.LibraryService
}

func New(s *infra.LibraryService) *Gateway {
	return &Gateway{
		service: s,
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

		res, err := g.service.CreateBook(r.Context(), &pb.CreateBookRequest{
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
			Limit  int64
			Offset int64
		}

		var req Req
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			JSONResponse(w, err, http.StatusInternalServerError)
			return
		}

		pb, err := g.service.ListBooks(r.Context(), &pb.ListBooksRequest{
			Limit:  req.Limit,
			Offset: req.Offset,
		})

		if err != nil {
			JSONResponse(w, err, http.StatusInternalServerError)
			return
		}

		books := mapper.ProtoToListBook(pb.Books)
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
