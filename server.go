package main

import (
	"fmt"
	"net/http"
	"strconv"
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

	a.Post("/api/todos", func(rw http.ResponseWriter, r *http.Request) {
		var t Todo
		Bind(r, &t)
		db = append(db, t)
		Json(rw, &db)
	})

	a.Delete("/api/todos", func(rw http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(QueryParam(r, "id"))
		if err != nil {
			panic(err)
		}
		var foundIdx int

		for i, todo := range db {
			if todo.Id == id {
				foundIdx = i
			}
		}

		if foundIdx > -1 {
			db = append(db[:foundIdx], db[foundIdx+1:]...)
		}

		Json(rw, &db)
	})

	a.Put("/api/todos", func(rw http.ResponseWriter, r *http.Request) {
		var t Todo
		Bind(r, &t)

		fmt.Printf("%+v\n", t)

		var foundIdx int

		for i, todo := range db {
			if todo.Id == t.Id {
				foundIdx = i
			}
		}

		if foundIdx > -1 {
			db[foundIdx].Completed = t.Completed
		}

		Json(rw, &db)
	})

	a.Static("examples/vanillajs")

	a.Start(":3000")
}
