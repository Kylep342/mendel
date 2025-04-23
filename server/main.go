package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	// db, err := sql.Open("mysql", "username:password@(127.0.0.1:3306)/dbname?parseTime=true")

	r := mux.NewRouter()

	// api := r.PathPrefix("/api")

	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "OK")
	})

	r.HandleFunc("/test/{thing}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		thing := vars["thing"]

		fmt.Fprintf(w, "TEST VALUE: %s \n", thing)
	})

	http.ListenAndServe(":8080", r)
}
