package response

type CibilResponse struct {
	Slug    string `json:"slug"`
	Score   string `json:"score"`
	Vendor  string `json:"vendor"`
	UserId  string `json:"user_id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Mobile  string `json:"mobile"`
	PanCard string `json:"pan_card"`
}
