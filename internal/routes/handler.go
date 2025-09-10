package router

import (
	"log"
	"net/http"

	"github.com/guruorgoru/adarsha-server/internal/models"
)

type SiteData struct {
	SiteName    string
	News        []models.NewsData
	Events      []models.EventData
	Vacancies   []models.VacanciesData
}

func rootHandler(templates *Templates) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := &SiteData{
			SiteName:    "Adarsha Secondary School",
			News:        models.GetDummyNews(),
			Events:      models.GetDummyEvents(),
			Vacancies:   models.GetDummyVacancies(),
		}
		log.Println("Attempting to render index.html")
		err := templates.Render(w, "index", data)
		if err != nil {
			log.Printf("Error rendering template: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func getNewsHandler(templates *Templates) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := &SiteData{
			SiteName:    "Adarsha Secondary School",
			News:        models.GetDummyNews(),
		}
		log.Println("Attempting to render news-page.html")
		err := templates.Render(w, "news-page", data)
		if err != nil {
			log.Printf("Error rendering template: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func getVacanciesHandler(templates *Templates) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := &SiteData{
			SiteName:    "Adarsha Secondary School",
			Vacancies:   models.GetDummyVacancies(),
		}
		log.Println("Attempting to render vacancies-page.html")
		err := templates.Render(w, "vacancies-page", data)
		if err != nil {
			log.Printf("Error rendering template: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func getEventsHandler(templates *Templates) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := &SiteData{
			SiteName:    "Adarsha Secondary School",
			Events:      models.GetDummyEvents(),
		}
		log.Println("Attempting to render events-page.html")
		err := templates.Render(w, "events-page", data)
		if err != nil {
			log.Printf("Error rendering template: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func createSimplePageHandler(templates *Templates, templateName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := &SiteData{
			SiteName: "Adarsha Secondary School",
		}
		log.Printf("Attempting to render %s.html", templateName)
		err := templates.Render(w, templateName, data)
		if err != nil {
			log.Printf("Error rendering template %s: %v", templateName, err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
