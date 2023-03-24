package types

//ProductCreateResponse to message for response handler products.
type ProductCreateResponse struct {
	Message string `json:"message,omitempty"`
}

/*
type ProductResponse struct {
	Message string         `json:"message"`
	Product *model.Product `json:"product,omitempty"`
	Error   string         `json:"error,omitempty"`
}

type ProductsResponse struct {
	Message string           `json:"message"`
	Product []*model.Product `json:"product,omitempty"`
	Error   string           `json:"error,omitempty"`
}*/
type GenericResponse struct {
	Message string      `json:"message"`
	Product interface{} `json:"product,omitempty"`
	Error   string      `json:"error,omitempty"`
}
