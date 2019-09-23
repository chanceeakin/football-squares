package common

// ID is a standard struct
type ID struct {
	ID *string `json:"id" validate:"required"`
}

// Success is a
type Success struct {
	Success bool `json:"success"`
}
