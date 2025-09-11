package db

import (
	"log"
	"os"

	supa "github.com/nedpals/supabase-go"
)

var SupabaseClient *supa.Client

func Init() {
	supabaseUrl := os.Getenv("DB_URL")
	supabaseKey := os.Getenv("DB_KEY")
	if supabaseUrl == "" || supabaseKey == "" {
		log.Fatalln("DB_URL or DB_KEY not set in environment")
	}
	SupabaseClient = supa.CreateClient(supabaseUrl, supabaseKey)
}
