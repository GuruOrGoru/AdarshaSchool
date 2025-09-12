package router

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

type Templates struct {
	templates *template.Template
}

func (t *Templates) Render(w http.ResponseWriter, name string, data any) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func NewTemplates() *Templates {
	tmpl := template.New("")
	tmpl = template.Must(template.ParseGlob("views/*.html"))
	tmpl = template.Must(tmpl.ParseGlob("views/academics/*.html"))
	tmpl = template.Must(tmpl.ParseGlob("views/athletics/*.html"))
	tmpl = template.Must(tmpl.ParseGlob("views/partials/*.html"))
	tmpl = template.Must(tmpl.ParseGlob("views/admissions/*.html"))
	return &Templates{
		templates: tmpl,
	}
}

func NewRouter(templates *Templates) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.StripSlashes)
	r.Use(middleware.Timeout(60 * 1e9))
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	r.Get("/", rootHandler(templates))
	r.Get("/news", getNewsHandler(templates))
	r.Get("/vacancies", getVacanciesHandler(templates))
	r.Get("/events", getEventsHandler(templates))
	r.Get("/login", loginHandler(templates))
	
	r.Group(func(r chi.Router) {
		r.Use(adminOnly)
		r.Get("/dashboard", dashboardHandler(templates))
	})

	r.Post("/news", postNewsHandler(templates))
	r.Post("/events", postEventHandler(templates))
	r.Post("/vacancies", postVacancyHandler(templates))
	r.Post("/login", loginHandler(templates))

	r.Get("/about", createSimplePageHandler(templates, "about"))
	r.Get("/academics", createSimplePageHandler(templates, "academics"))
	r.Get("/academics/graduate", createSimplePageHandler(templates, "academics/graduate"))
	r.Get("/academics/research", createSimplePageHandler(templates, "academics/research"))
	r.Get("/academics/schools", createSimplePageHandler(templates, "academics/schools"))
	r.Get("/academics/undergraduate", createSimplePageHandler(templates, "academics/undergraduate"))
	r.Get("/admissions", createSimplePageHandler(templates, "admissions"))
	r.Get("/admissions/financial-aid", createSimplePageHandler(templates, "admissions/financial-aid"))
	r.Get("/admissions/graduate", createSimplePageHandler(templates, "admissions/graduate"))
	r.Get("/admissions/undergraduate", createSimplePageHandler(templates, "admissions/undergraduate"))
	r.Get("/admissions/visit", createSimplePageHandler(templates, "admissions/visit"))
	r.Get("/affiliates", createSimplePageHandler(templates, "affiliates"))
	r.Get("/affordability", createSimplePageHandler(templates, "affordability"))
	r.Get("/apply", createSimplePageHandler(templates, "apply"))
	r.Get("/athletics", createSimplePageHandler(templates, "athletics"))
	r.Get("/athletics/all-sports", createSimplePageHandler(templates, "athletics/all-sports"))
	r.Get("/athletics/basketball", createSimplePageHandler(templates, "athletics/basketball"))
	r.Get("/athletics/football", createSimplePageHandler(templates, "athletics/football"))
	r.Get("/athletics/tickets", createSimplePageHandler(templates, "athletics/tickets"))
	r.Get("/campus-life", createSimplePageHandler(templates, "campus-life"))
	r.Get("/campus-map", createSimplePageHandler(templates, "campus-map"))
	r.Get("/contact", createSimplePageHandler(templates, "contact"))
	r.Get("/directories", createSimplePageHandler(templates, "directories"))
	r.Get("/downtown", createSimplePageHandler(templates, "downtown"))
	r.Get("/emergency", createSimplePageHandler(templates, "emergency"))
	r.Get("/ethics", createSimplePageHandler(templates, "ethics"))
	r.Get("/experience", createSimplePageHandler(templates, "experience"))
	r.Get("/give", createSimplePageHandler(templates, "give"))
	r.Get("/giving-societies", createSimplePageHandler(templates, "giving-societies"))
	r.Get("/history", createSimplePageHandler(templates, "history"))
	r.Get("/hotline", createSimplePageHandler(templates, "hotline"))
	r.Get("/hr", createSimplePageHandler(templates, "hr"))
	r.Get("/libraries", createSimplePageHandler(templates, "libraries"))
	r.Get("/non-discrimination", createSimplePageHandler(templates, "non-discrimination"))
	r.Get("/privacy", createSimplePageHandler(templates, "privacy"))
	r.Get("/schools", createSimplePageHandler(templates, "schools"))
	r.Get("/search", createSimplePageHandler(templates, "search"))
	r.Get("/student-life", createSimplePageHandler(templates, "student-life"))
	r.Get("/visitor-info", createSimplePageHandler(templates, "visitor-info"))
	r.Get("/wellbeing", createSimplePageHandler(templates, "wellbeing"))

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	staticDir := filepath.Join(cwd, "static")
	log.Println("Serving static files from:", staticDir)
	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir(staticDir))))

	uploadsDir := filepath.Join(cwd, "uploads")
	log.Println("Serving uploaded files from:", uploadsDir)
	r.Handle("/uploads/*", http.StripPrefix("/uploads/", http.FileServer(http.Dir(uploadsDir))))

	return r
}
