package main

import (
	"encoding/json"
	"net/http"
)

type Todo struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Complted bool   `json:"completed"`
}

func main() {
	db := []Todo{}

	a := NewApplication()

	a.Get("/api/todos", func(rw http.ResponseWriter, r *http.Request) {
		enc := json.NewEncoder(rw)
		enc.Encode(&db)

		// res.Json(&db)
	})

	a.Post("/api/addTodo", func(rw http.ResponseWriter, r *http.Request) {
		var t Todo

		// req.Bind(&t)

		json.NewDecoder(r.Body).Decode(&t)
		db = append(db, t)

		enc := json.NewEncoder(rw)
		enc.Encode(&db)

		// res.Json(&db)
	})

	a.Static("examples/vanillajs")

	a.Start(":3000")
}
