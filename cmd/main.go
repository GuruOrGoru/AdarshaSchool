package main

import (
	"log"
	"net/http"
	"os"

	"github.com/guruorgoru/adarsha-server/internal/db"
	router "github.com/guruorgoru/adarsha-server/internal/routes"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file, proceeding with system environment variables")
	}
	db.Init()
	port := os.Getenv("PORT")
	if port == ""{
		log.Fatalln("PORT not set in environment")
	}
	templates := router.NewTemplates()
	r := router.NewRouter(templates)

	log.Println("Server started at port:", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
