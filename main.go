package main

import (
	"fmt"
	"log"
	"net/http"

	"example.com/myapi/handler"
	"example.com/myapi/middleware"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/user", handler.UserHandler)

	// CORSミドルウェアでラップ
	handlerWithCORS := middleware.CORS(mux)

	fmt.Println("🚀 Running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", handlerWithCORS))
}
