package router

import (
	"net/http"
	"text/template"

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
	tmpl = template.Must(tmpl.ParseGlob("views/*.html"))
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
	r.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	return r
}
