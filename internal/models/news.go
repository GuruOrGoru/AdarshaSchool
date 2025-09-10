package models

type NewsData struct {
	Title       string
	Description string
	ImageURL    string
}

func GetDummyNews() []NewsData {
	return []NewsData{
		{
			Title:       "School Reopens for New Academic Year",
			Description: "Adarsha Secondary School is excited to welcome students back for the new academic year with enhanced safety measures.",
			ImageURL:    "/static/logo.jpg",
		},
		{
			Title:       "Annual Sports Day Highlights",
			Description: "The Annual Sports Day was a grand success with students showcasing their athletic talents and team spirit.",
			ImageURL:    "/static/logo.jpg",
		},
		{
			Title:       "Science Fair Winners Announced",
			Description: "Congratulations to the winners of the Science Fair! Their innovative projects impressed both judges and attendees.",
			ImageURL:    "/static/logo.jpg",
		},
		{
			Title:       "New Library Inauguration",
			Description: "The new library at Adarsha Secondary School has been inaugurated, providing students with a vast collection of books and resources.",
			ImageURL:    "/static/logo.jpg",
		},
		{
			Title:       "Cultural Fest Celebrates Diversity",
			Description: "The Cultural Fest was a vibrant celebration of the diverse cultures represented at our school, featuring performances, food stalls, and art exhibitions.",
			ImageURL:    "/static/logo.jpg",
		},
	}
}

