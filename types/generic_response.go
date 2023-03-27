package types

//ProductCreateResponse to message for response handler products.
type CreateResponse struct {
	Message string `json:"message,omitempty"`
}

type GenericResponse struct {
	Message string      `json:"message"`
	Product interface{} `json:"product,omitempty"`
	Error   string      `json:"error,omitempty"`
}

type GenericUserResponse struct {
	Message string      `json:"message"`
	User    interface{} `json:"product,omitempty"`
	Error   string      `json:"error,omitempty"`
}
