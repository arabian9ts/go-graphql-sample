package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		result := executeUserQuery(r.URL.Query().Get("query"), userSchema)
		json.NewEncoder(w).Encode(result)
	})

	http.HandleFunc("/message", func(w http.ResponseWriter, r *http.Request) {
		result := executeMessageQuery(r.URL.Query().Get("query"), messageSchema)
		json.NewEncoder(w).Encode(result)
	})

	fmt.Println("server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
