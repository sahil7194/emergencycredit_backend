package request

type CreateCityRequest struct {
	Name    string `json:"name" binding:"required"`
	StateId string `json:"state_id" binding:"required"`
}
