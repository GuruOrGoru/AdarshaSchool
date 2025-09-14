package models

import (
	"github.com/guruorgoru/adarsha-server/internal/db"
)

type EventData struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Date        string `json:"date"`
	ImageURL    string `json:"image_url"`
}

func InsertEvent(event EventData) ([]EventData, error) {
	var results []EventData
	err := db.SupabaseClient.DB.From("events").Insert(event).Execute(&results)
	return results, err
}

func GetAllEvent() ([]EventData, error) {
	var results []EventData
	err := db.SupabaseClient.DB.From("events").Select("*").Execute(&results)
	return results, err
}

func SearchEvents(query string) ([]EventData, error) {
	var results []EventData
	err := db.SupabaseClient.DB.From("events").
		Select("*").
		Filter("title", "ilike", "%"+query+"%").
		Execute(&results)
	return results, err
}

func UpdateEvent(id string, news EventData) ([]EventData, error) {
	var results []EventData
	err := db.SupabaseClient.DB.From("events").Update(news).Eq("id", id).Execute(&results)
	return results, err
}

func DeleteEvent(id string) error {
	err := db.SupabaseClient.DB.From("events").Delete().Eq("id", id).Execute(nil)
	return err
}
