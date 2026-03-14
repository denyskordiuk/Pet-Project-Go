package main

import (
	"fmt"
	"net/http"
	fucntionsTool "petproject/tools"
)

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/hello", fucntionsTool.HelloHandler)

	router.HandleFunc("GET /users/{id}", fucntionsTool.GetUser)
	router.HandleFunc("POST /users", fucntionsTool.CreateUser)
	router.HandleFunc("DELETE /users/{id}", fucntionsTool.DeleteUser)

	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		fmt.Println("Error starting server:", err)
		panic(err)
	}
}
