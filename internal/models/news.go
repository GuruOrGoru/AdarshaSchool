package models

import "github.com/guruorgoru/adarsha-server/internal/db"

type NewsData struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ImageURL    string `json:"image_url"`
}

func InsertNews(news NewsData) ([]NewsData, error) {
	var results []NewsData
	err := db.SupabaseClient.DB.From("news").Insert(news).Execute(&results)
	return results, err
}

func GetAllNews() ([]NewsData, error) {
	var results []NewsData
	err := db.SupabaseClient.DB.From("news").Select("*").Execute(&results)
	return results, err
}

func UpdateNews(id string, news NewsData) ([]NewsData, error) {
	var results []NewsData
	err := db.SupabaseClient.DB.From("news").Update(news).Eq("id", id).Execute(&results)
	return results, err
}

func DeleteNews(id string) error {
	err := db.SupabaseClient.DB.From("news").Delete().Eq("id", id).Execute(nil)
	return err
}

func SearchNews(query string) ([]NewsData, error) {
	var results []NewsData
	err := db.SupabaseClient.DB.From("news").
		Select("*").
		Filter("title", "ilike", "%"+query+"%").
		Execute(&results)
	return results, err
}
