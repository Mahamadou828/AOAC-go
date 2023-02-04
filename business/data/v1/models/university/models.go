package university

import "time"

// University represents an available university that user can apply.
// Universities are obtain using scrapper that scrape the website
// https://www.hotcoursesabroad.com/study/international/schools-colleges-university/list.html?sortby=ALL&pageNo=%d
type University struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Country     string    `json:"country"`
	DetailsURL  string    `json:"detailsURL"`
	Description string    `json:"description"`
	Faculties   []string  `json:"faculties"`
	Rating      int       `json:"logoURL"`
	CreatedAt   time.Time `json:"createdAt"`
	DeleteAt    time.Time `json:"deleteAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
