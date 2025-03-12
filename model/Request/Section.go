package request


type ReqSection struct {

	Name        string `json:"name" validate:"required,max=10"`    
	Description string `json:"description" validate:"max=100"` 


}