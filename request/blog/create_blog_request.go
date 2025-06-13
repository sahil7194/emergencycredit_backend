package request

type CreateBlogRequest struct {
	Title   string `json:"title" binding:"required"`
	Summary string `json:"summary" binding:"required"`
	Content string `json:"content" binding:"required"`
}
