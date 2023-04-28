package domain

type TotalsByCustomerCondition struct {
	Inactive float64 `json:"inactive"`
	Active   float64 `json:"active"`
}

type TopSoldProduct struct {
	Description string `json:"description"`
	Total       int    `json:"total"`
}

type TopActiveCustomerSpent struct {
	LastName  string  `json:"last_name"`
	FirstName string  `json:"first_name"`
	Amount    float64 `json:"amount"`
}
