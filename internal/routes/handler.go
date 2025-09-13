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
	SiteName    string
	News        []models.NewsData
	Events      []models.EventData
	Vacancies   []models.VacanciesData
	IsAdmin     bool
	CurrentPage string
}

func rootHandler(templates *Templates) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		FromModel, err := models.GetAllNews()
		if err != nil {
			log.Printf("Error during  retrieval: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		eventsFromModel, err := models.GetAllEvent()
		if err != nil {
			log.Printf("Error during  retrival: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		vacanciesFromModel, err := models.GetAllVacancies()
		if err != nil {
			log.Printf("Error during  retrival: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		data := &SiteData{
			SiteName:    "Adarsha Secondary School",
			News:        FromModel,
			Events:      eventsFromModel,
			Vacancies:   vacanciesFromModel,
			IsAdmin:     models.IsLoggedIn(r),
			CurrentPage: "home",
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
		FromModel, err := models.GetAllNews()
		if err != nil {
			log.Printf("Error during  retrieval: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		data := &SiteData{
			SiteName:    "Adarsha Secondary School",
			News:        FromModel,
			CurrentPage: "news",
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
		vacanciesFromModel, err := models.GetAllVacancies()
		if err != nil {
			log.Printf("Error while retrieval: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		data := &SiteData{
			SiteName:    "Adarsha Secondary School",
			Vacancies:   vacanciesFromModel,
			CurrentPage: "vacancies",
		}
		log.Println("Attempting to render vacancies-page.html")
		err = templates.Render(w, "vacancies-page", data)
		if err != nil {
			log.Printf("Error rendering template: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func getEventsHandler(templates *Templates) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		eventsFromModel, err := models.GetAllEvent()
		if err != nil {
			log.Printf("Error during  retrival: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		data := &SiteData{
			SiteName:    "Adarsha Secondary School",
			Events:      eventsFromModel,
			CurrentPage: "events",
		}
		log.Println("Attempting to render events-page.html")
		err = templates.Render(w, "events-page", data)
		if err != nil {
			log.Printf("Error rendering template: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func createSimplePageHandler(templates *Templates, templateName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := &SiteData{
			SiteName:    "Adarsha Secondary School",
			CurrentPage: templateName,
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
		var imageUrl string
		if err != nil {
			imageUrl = ""
		} else {
			defer file.Close()
			err = os.MkdirAll("./uploads", os.ModePerm)
			if err != nil {
				log.Printf("Error creating uploads directory: %v", err)
				http.Error(w, "Internal server error", 500)
				return
			}

			rawFileName := "news-" + header.Filename

			filename := filepath.Join("uploads", rawFileName)
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
			imageUrl = "/" + filename
		}

		Item := models.NewsData{
			Id:          rand.Intn(900000000) + 100000000,
			Title:       title,
			Description: description,
			ImageURL:    imageUrl,
		}

		List, err := models.InsertNews(Item)
		if err != nil {
			http.Error(w, "Failed to insert : "+err.Error(), http.StatusInternalServerError)
			return
		}

		data := &SiteData{
			News: List,
		}

		err = templates.Render(w, "news-partial", data)
		if err != nil {
			log.Printf("Error rendering -partial template: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		log.Printf("News item created: %+v", Item)
	}
}

func postEventHandler(templates *Templates) http.HandlerFunc {
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
		dateStr := r.FormValue("date")

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

		rawFileName := "event-" + header.Filename

		filename := filepath.Join("uploads", rawFileName)
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

		eventItem := models.EventData{
			Id:          rand.Intn(900000000) + 100000000,
			Title:       title,
			Description: description,
			Date:        dateStr,
			ImageURL:    "/" + filename,
		}

		eventList, err := models.InsertEvent(eventItem)
		if err != nil {
			http.Error(w, "Failed to insert event: "+err.Error(), http.StatusInternalServerError)
			return
		}

		data := &SiteData{
			Events: eventList,
		}

		err = templates.Render(w, "events-partial", data)
		if err != nil {
			log.Printf("Error rendering events template: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		log.Printf("Events item created: %+v", eventItem)
	}
}

func postVacancyHandler(templates *Templates) http.HandlerFunc {
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
		var imageUrl string
		if err != nil {
			imageUrl = ""
		} else {
			defer file.Close()
			err = os.MkdirAll("./uploads", os.ModePerm)
			if err != nil {
				log.Printf("Error creating uploads directory: %v", err)
				http.Error(w, "Internal server error", 500)
				return
			}

			rawFileName := "vacancies-" + header.Filename

			filename := filepath.Join("uploads", rawFileName)
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
			imageUrl = "/" + filename
		}

		Item := models.VacanciesData{
			Id:          rand.Intn(900000000) + 100000000,
			Title:       title,
			Description: description,
			ImageURL:    imageUrl,
		}

		List, err := models.InsertVacancies(Item)
		if err != nil {
			http.Error(w, "Failed to insert : "+err.Error(), http.StatusInternalServerError)
			return
		}

		data := &SiteData{
			Vacancies: List,
		}

		err = templates.Render(w, "vacancies-partial", data)
		if err != nil {
			log.Printf("Error rendering vacancies-partial template: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		log.Printf("Vacancies item created: %+v", Item)
	}
}

func loginHandler(templates *Templates) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			data := &SiteData{
				SiteName:    "Adarsha Secondary School",
				CurrentPage: "login",
			}
			err := templates.Render(w, "login", data)
			if err != nil {
				log.Printf("Error rendering login template: %v", err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		case http.MethodPost:
			err := r.ParseForm()
			if err != nil {
				http.Error(w, "Failed to parse form: "+err.Error(), http.StatusBadRequest)
				return
			}
			email := r.FormValue("email")
			password := r.FormValue("password")

			if models.IsAdmin(email, password) {
				models.SetSession(w)
				http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
				return
			} else {
				http.Error(w, "Invalid credentials", http.StatusUnauthorized)
				return
			}
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
	}
}

func dashboardHandler(templates *Templates) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := &SiteData{
			SiteName:    "Adarsha Secondary School",
			IsAdmin:     true,
			CurrentPage: "dashboard",
		}
		err := templates.Render(w, "dashboard", data)
		if err != nil {
			log.Printf("Error rendering dashboard template: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	models.ClearSession(w)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func adminOnly(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !models.IsLoggedIn(r) {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		next.ServeHTTP(w, r)
	})
}
