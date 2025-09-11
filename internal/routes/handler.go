package router

import (
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"

	"github.com/guruorgoru/adarsha-server/internal/models"
)

type SiteData struct {
	SiteName  string
	News      []models.NewsData
	Events    []models.EventData
	Vacancies []models.VacanciesData
}

func rootHandler(templates *Templates) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		newsFromModel, err := models.GetAllNews()
		if err != nil {
			log.Printf("Error during news retrieval: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		data := &SiteData{
			SiteName:  "Adarsha Secondary School",
			News:      newsFromModel,
			Events:    models.GetDummyEvents(),
			Vacancies: models.GetDummyVacancies(),
		}
		log.Println("Attempting to render index.html")
		err = templates.Render(w, "index", data)
		if err != nil {
			log.Printf("Error rendering template: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func getNewsHandler(templates *Templates) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		newsFromModel, err := models.GetAllNews()
		if err != nil {
			log.Printf("Error during news retrieval: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		data := &SiteData{
			SiteName: "Adarsha Secondary School",
			News:     newsFromModel,
		}
		log.Println("Attempting to render news-page.html")
		err = templates.Render(w, "news-page", data)
		if err != nil {
			log.Printf("Error rendering template: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func getVacanciesHandler(templates *Templates) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := &SiteData{
			SiteName:  "Adarsha Secondary School",
			Vacancies: models.GetDummyVacancies(),
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
			SiteName: "Adarsha Secondary School",
			Events:   models.GetDummyEvents(),
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

func postNewsHandler(templates *Templates) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		err := r.ParseMultipartForm(10 << 20)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		title := r.FormValue("title")
		description := r.FormValue("description")

		file, header, err := r.FormFile("image")
		if err != nil {
			http.Error(w, "Image upload required", http.StatusBadRequest)
			return
		}
		defer file.Close()
		err = os.MkdirAll("./uploads", os.ModePerm)
		if err != nil {
			log.Printf("Error creating uploads directory: %v", err)
			http.Error(w, "Internal server error", 500)
			return
		}

		filename := filepath.Join("uploads", header.Filename)
		dst, err := os.Create(filename)
		if err != nil {
			http.Error(w, "Error saving file: "+err.Error(), 500)
			return
		}
		defer dst.Close()
		_, err = io.Copy(dst, file)
		if err != nil {
			http.Error(w, "Error saving file: "+err.Error(), 500)
			return
		}

		newsItem := models.NewsData{
			Id:          rand.Intn(900000000) + 100000000,
			Title:       title,
			Description: description,
			ImageURL:    "/" + filename,
		}

		newsList, err := models.InsertNews(newsItem)
		if err != nil {
			http.Error(w, "Failed to insert news: "+err.Error(), http.StatusInternalServerError)
			return
		}

		data := &SiteData{
			News:     newsList,
		}

		err = templates.Render(w, "news-partial", data)
		if err != nil {
			log.Printf("Error rendering news-partial template: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		log.Printf("News item created: %+v", newsItem)
	}
}
