package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Msg string `json:"msg"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Msg: "success",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/api/hello", helloHandler)

	port := 8080
	fmt.Printf("Server started on port %d\n", port)
	addr := fmt.Sprintf(":%d", port)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
