package request

type CheckCibil struct {
	Slug string `json:"slug" binding:"required"`
}
