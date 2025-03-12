package request

type CreatePostRequest struct {
	Title     string `json:"title" validate:"required,min=1,max=200"`
	Content   string `json:"content" validate:"required,min=1"`
	SectionID uint   `json:"section_id" validate:"required"`
}
