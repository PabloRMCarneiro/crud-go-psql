package main

import (
	"net/http"
	"crud/configs"
	"crud/handlers"
	"github.com/go-chi/chi/v5"
)

func main() {

	err := configs.Load()
	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()

	r.Post("/", handlers.Create)
	r.Put("/{id}", handlers.Update)
	r.Delete("/{id}", handlers.Delete)
	r.Get("/{id}", handlers.Get)
	r.Get("/", handlers.List)

	http.ListenAndServe(":9000", r)


}