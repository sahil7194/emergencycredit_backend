package request

type UpdateSchemeRequest struct {
	Title           string `json:"title" binding:"required"`
	Summary         string `json:"summary" binding:"required"`
	Description     string `json:"description" binding:"required"`
	MaxAmount       string `json:"max_amount" binding:"required"`
	MinAmount       string `json:"min_amount" binding:"required"`
	MaxCibil        string `json:"max_cibil" binding:"required"`
	MinCibil        string `json:"min_cibil" binding:"required"`
	MaxInterestRate string `json:"max_interest_rate" binding:"required"`
	MinInterestRate string `json:"min_interest_rate" binding:"required"`
	MaxTenure       string `json:"max_tenure" binding:"required"`
	MinTenure       string `json:"min_tenure" binding:"required"`
	RedirectionLink string `json:"redirection_link" binding:"required"`
	BankId          string `json:"bank_id" binding:"required"`
}
