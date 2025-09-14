package models

import (
	"github.com/guruorgoru/adarsha-server/internal/db"
)

type StaffData struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Position string `json:"position"`
	ImageURL string `json:"image_url"`
}

func InsertStaff(staff StaffData) ([]StaffData, error) {
	var results []StaffData
	err := db.SupabaseClient.DB.From("staffs").Insert(staff).Execute(&results)
	return results, err
}

func GetAllStaff() ([]StaffData, error) {
	var results []StaffData
	err := db.SupabaseClient.DB.From("staffs").Select("*").Execute(&results)
	return results, err
}

func SearchStaff(query string) ([]StaffData, error) {
	var results []StaffData
	err := db.SupabaseClient.DB.From("staffs").
		Select("*").
		Filter("name", "ilike", "%"+query+"%").
		Execute(&results)
	return results, err
}

func UpdateStaffs(id string, staffs StaffData) ([]StaffData, error) {
	var results []StaffData
	err := db.SupabaseClient.DB.From("staffs").Update(staffs).Eq("id", id).Execute(&results)
	return results, err
}

func DeleteStaffs(id string) error {
	err := db.SupabaseClient.DB.From("staffs").Delete().Eq("id", id).Execute(nil)
	return err
}
