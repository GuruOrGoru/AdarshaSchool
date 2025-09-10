package models

type VacanciesData struct {
	Title       string
	Description string
	ImageURL    string
}

func GetDummyVacancies() []VacanciesData {
	return []VacanciesData{
		{
			Title:       "Software Engineer",
			Description: "Join our team as a Software Engineer to develop innovative solutions.",
			ImageURL:    "https://example.com/images/software-engineer.jpg",
		},		{
			Title:       "Product Manager",
			Description: "We are looking for a Product Manager to lead product development.",
			ImageURL:    "https://example.com/images/product-manager.jpg",
		},
		{
			Title:       "UX Designer",
			Description: "Seeking a creative UX Designer to enhance user experiences.",
			ImageURL:    "https://example.com/images/ux-designer.jpg",
		},
		{
			Title:       "Data Scientist",
			Description: "Looking for a Data Scientist to analyze and interpret complex data.",
			ImageURL:    "https://example.com/images/data-scientist.jpg",
		},
		{
			Title:       "Marketing Specialist",
			Description: "Join us as a Marketing Specialist to drive our marketing strategies.",
			ImageURL:    "https://example.com/images/marketing-specialist.jpg",
		},
	}
}
