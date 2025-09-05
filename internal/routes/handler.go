package router

import (
	"log"
	"net/http"
)

type NavBarData struct {
	SiteName string
}

func rootHandler(templates *Templates) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := &NavBarData{
			SiteName: "Adarsha Secondary School",
		}
		log.Println("Attempting to render index.html")
		err := templates.Render(w, "index.html", data)
		if err != nil {
			log.Printf("Error rendering template: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
