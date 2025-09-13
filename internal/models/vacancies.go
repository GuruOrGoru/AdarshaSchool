package models

import "github.com/guruorgoru/adarsha-server/internal/db"

type VacanciesData struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ImageURL    string `json:"image_url"`
}

func InsertVacancies(vacancies VacanciesData) ([]VacanciesData, error) {
	var results []VacanciesData
	err := db.SupabaseClient.DB.From("vacancies").Insert(vacancies).Execute(&results)
	return results, err
}

func GetAllVacancies() ([]VacanciesData, error) {
	var results []VacanciesData
	err := db.SupabaseClient.DB.From("vacancies").Select("*").Execute(&results)
	return results, err
}

func SearchVacancies(query string) ([]VacanciesData, error) {
	var results []VacanciesData
	err := db.SupabaseClient.DB.From("vacancies").
		Select("*").
		Filter("title", "ilike", "%"+query+"%").
		Execute(&results)
	return results, err
}
