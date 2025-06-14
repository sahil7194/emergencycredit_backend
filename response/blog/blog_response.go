package response

type BlogResponse struct {
	Title   string `json:"title"`
	Slug    string `json:"slug"`
	Summary string `json:"summary"`
	Content string `json:"content"`
	Image   string `json:"image" `
	Status  string `json:"status"`
}
