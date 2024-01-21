package dtos

type CreateCategoryDto struct {
	CategoryType string `json:"category_type,omitempty" binding:"required,currency" oneof="Hot New Top Rising"`
	Name         string `json:"name"`
}
