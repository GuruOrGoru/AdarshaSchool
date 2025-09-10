package models

import "time"

type EventData struct {
	Title       string
	Description string
	Date        time.Time
	ImageURL    string
}

func GetDummyEvents() []EventData {
	return []EventData{
		{
			Title:       "Annual Sports Day",
			Description: "Join us for a day of fun and competition as students showcase their athletic talents.",
			Date:        time.Date(2023, 5, 20, 9, 0, 0, 0, time.UTC),
			ImageURL:    "/static/sports_day.jpg",
		},
		{
			Title:       "Science Fair",
			Description: "Explore innovative projects and experiments presented by our budding scientists.",
			Date:        time.Date(2023, 6, 15, 10, 0, 0, 0, time.UTC),
			ImageURL:    "/static/science_fair.jpg",
		},
		{
			Title:       "Cultural Fest",
			Description: "Celebrate the rich cultural diversity of our school with performances, food, and art.",
			Date:        time.Date(2023, 7, 10, 11, 0, 0, 0, time.UTC),
			ImageURL:    "/static/cultural_fest.jpg",
		},
		{
			Title:       "Library Inauguration",
			Description: "Join us for the grand opening of our new library, a hub for learning and exploration.",
			Date:        time.Date(2023, 8, 5, 14, 0, 0, 0, time.UTC),
			ImageURL:    "/static/library_inauguration.jpg",
		},
		{
			Title:       "Art Exhibition",
			Description: "Experience the creativity of our students through a diverse range of artworks on display.",
			Date:        time.Date(2023, 9, 12, 13, 0, 0, 0, time.UTC),
			ImageURL:    "/static/art_exhibition.jpg",
		},
		{
			Title:       "Music Concert",
			Description: "Enjoy an evening of melodious performances by our talented student musicians.",
			Date:        time.Date(2023, 10, 18, 18, 0, 0, 0, time.UTC),
			ImageURL:    "/static/music_concert.jpg",
		},
	}
}
