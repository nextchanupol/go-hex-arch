package dto

type CustomerResponse struct {
	CustomerID  int64  `json:"customer_id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	DateOfBirth string `json:"date_of_birth"`
	Email       string `json:"email"`
	IsActive    bool   `json:"is_active"`
	NoOfActive  int    `json:"no_of_active"`
}
