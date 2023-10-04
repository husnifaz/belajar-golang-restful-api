package web

type StudentCreateRequest struct {
	Name string `validate:"required,min=1,max=100" json:"name"`
	Address string `validate:"required" json:"address"`
}
