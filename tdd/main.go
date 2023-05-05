package main

import (
	"log"
	"net/http"
	"os"

	"github.com/addetz/gophertale/tdd/data"
	"github.com/addetz/gophertale/tdd/handlers"
)

func main() {
	es := data.NewEmployeeService()
	ha := handlers.NewHandler(es)
	http.HandleFunc("/employees", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			ha.ListEmployeeByID(w, r)
		case "POST":
			ha.UpsertEmployee(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server listening on port %s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
