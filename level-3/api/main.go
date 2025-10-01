package main

import (
    "fmt"
    "log"
    "net/http"

    "github.com/gorilla/mux"
)

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/account/{id}", getAccount).Methods("GET")

    fmt.Println("Server running on port 8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}

func getAccount(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    // TODO: Connect to Fabric network and fetch account details
    fmt.Fprintf(w, "Account ID: %s", id)
}
