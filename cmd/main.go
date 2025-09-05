package main

import (
	"log"
	"net/http"

	router "github.com/guruorgoru/adarsha-server/internal/routes"
)

func main() {
	templates := router.NewTemplates()
	r := router.NewRouter(templates)

	log.Println("Server started at port: 8414")
	log.Fatal(http.ListenAndServe(":8414", r))
}
