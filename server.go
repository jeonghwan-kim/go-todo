package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Todo struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Complted bool   `json:"completed"`
}

func main() {
	db := []Todo{}

	http.HandleFunc("/api/todos", func(rw http.ResponseWriter, r *http.Request) {
		// t := Todo{1, "todo1", false}
		// db = append(db, t)

		enc := json.NewEncoder(rw)
		enc.Encode(&db)
	})

	http.HandleFunc("/api/addTodo", func(rw http.ResponseWriter, r *http.Request) {
		var t Todo

		json.NewDecoder(r.Body).Decode(&t)
		db = append(db, t)

		enc := json.NewEncoder(rw)
		enc.Encode(&db)
	})

	http.Handle("/", http.FileServer(http.Dir("examples/vanillajs")))

	fmt.Println("server is running http://localhost:3000")
	http.ListenAndServe(":3000", nil)
}
