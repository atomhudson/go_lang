package types

type Student struct {
	ID    int64
	Name  string `validate:"required"`
	Email string `validate:"required,email"`
	Phone string `validate:"required"`
	Age   int    `validate:"required"`
}

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}
