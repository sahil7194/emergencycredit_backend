package request

type CreateStateRequest struct {
	Name string `json:"name" binding:"required"`
}
