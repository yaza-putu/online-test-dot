package validation

type Goods struct {
	Name       string `form:"name" json:"name" validate:"required"`
	CategoryId string `form:"category_id" json:"category_id" validate:"required"`
}
