package response

type SchemeResponse struct {
	Summary         string `json:"summary"`
	Slug            string `json:"slug" `
	Title           string `json:"title"`
	Description     string `json:"description"`
	Image           string `json:"image,omitempty"`
	RedirectionLink string `json:"redirection_link"`
	MinInterestRate string `json:"min_interest_rate"`
	MaxInterestRate string `json:"max_interest_rate"`
	MinCIBIL        string `json:"min_cibil"`
	MaxCIBIL        string `json:"max_cibil"`
	MinTenure       string `json:"min_tenure"`
	MaxTenure       string `json:"max_tenure"`
	MinAmount       string `json:"min_amount"`
	MaxAmount       string `json:"max_amount"`
	Status          string `json:"status"`
}
