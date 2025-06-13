package request

type UpdateStateRequest struct {
	Name string `json:"name" binding:"required"`
}
