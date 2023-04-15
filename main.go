package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const PORT = ":7777"

func Hello(w http.ResponseWriter, r *http.Request) {
	msg := map[string]interface{}{
		"message": "halo",
	}

	payload, _ := json.Marshal(msg)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(payload)
}
func main() {

	r := mux.NewRouter()
	r.HandleFunc("/index", Hello).Methods("GET")

	fmt.Println("running on port", PORT)
	// http.ListenAndServe(PORT, r)
	err := http.ListenAndServeTLS(PORT, "server.crt", "server.key", r)
	if err != nil {
		log.Fatal("listen and serve:", err)
	}
}
