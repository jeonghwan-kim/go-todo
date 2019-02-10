package main

import (
	"fmt"
	"net/http"
)

type Todo struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	db := []Todo{}

	a := NewApplication()

	a.Get("/api/todos", func(rw http.ResponseWriter, r *http.Request) {
		Json(rw, &db)
	})

	a.Post("/api/addTodo", func(rw http.ResponseWriter, r *http.Request) {
		var t Todo
		// json.NewDecoder(r.Body).Decode(&t)
		Bind(r, &t)
		fmt.Printf("%+v\n", t)
		db = append(db, t)
		fmt.Printf("%+v\n", db)
		Json(rw, &db)
	})

	a.Static("examples/vanillajs")

	a.Start(":3000")
}
