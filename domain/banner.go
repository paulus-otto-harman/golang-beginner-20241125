package domain

type Banner struct {
	Photo    string `json:"photo"`
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
	PathPage string `json:"path_page"`
}
