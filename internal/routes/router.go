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
	tmpl = template.Must(tmpl.ParseGlob("views/partials/*.html"))
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
	r.Get("/logout", logoutHandler)
	
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
	r.Get("/admissions", createSimplePageHandler(templates, "admissions"))
	r.Get("/athletics", createSimplePageHandler(templates, "athletics"))
	r.Get("/contact", createSimplePageHandler(templates, "contact"))
	r.Get("/facilities", createSimplePageHandler(templates, "facilities"))
	r.Get("/privacy", createSimplePageHandler(templates, "privacy"))
	r.Get("/team", createSimplePageHandler(templates, "team"))
	r.Get("/search", createSimplePageHandler(templates, "search"))

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
