package main

import (
	"github.com/Gabo-div/bingo/apps/backend-main/internal/echo"
	"github.com/Gabo-div/bingo/apps/backend-main/internal/user"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	echo.Register(r)
	user.Register(r)

	log.Printf("Running on localhost:3001")

	http.ListenAndServe(
		"localhost:3001",
		r,
	)
}
