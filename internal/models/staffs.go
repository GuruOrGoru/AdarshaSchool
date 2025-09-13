package models

import (
	"github.com/guruorgoru/adarsha-server/internal/db"
)

type StaffData struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Role     string `json:"role"`
	ImageURL string `json:"image_url"`
}

func InsertStaff(staff StaffData) ([]StaffData, error) {
	var results []StaffData
	err := db.SupabaseClient.DB.From("Staffs").Insert(staff).Execute(&results)
	return results, err
}

func GetAllStaff() ([]StaffData, error) {
	var results []StaffData
	err := db.SupabaseClient.DB.From("Staffs").Select("*").Execute(&results)
	return results, err
}

func SearchStaff(query string) ([]StaffData, error) {
	var results []StaffData
	err := db.SupabaseClient.DB.From("Staffs").
		Select("*").
		Filter("name", "ilike", "%"+query+"%").
		Execute(&results)
	return results, err
}
