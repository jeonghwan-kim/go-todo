package main

import (
	"net/http"
	"strconv"
)

func main() {
	a := NewApplication()
	s := NewStore()

	a.Get("/api/todos", func(rw http.ResponseWriter, r *http.Request) {
		Json(rw, s.FindAll())
	})

	a.Post("/api/todos", func(rw http.ResponseWriter, r *http.Request) {
		var t Todo
		Bind(r, &t)
		s.Create(t)
		Json(rw, s.FindAll())
	})

	a.Delete("/api/todos", func(rw http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(QueryParam(r, "id"))
		if err != nil {
			panic(err)
		}
		s.Destroy(id)
		Json(rw, s.FindAll())
	})

	a.Put("/api/todos", func(rw http.ResponseWriter, r *http.Request) {
		var t Todo
		Bind(r, &t)
		s.Update(t)
		Json(rw, s.FindAll())
	})

	a.Static("examples/vanillajs")

	a.Start(":3000")
}
