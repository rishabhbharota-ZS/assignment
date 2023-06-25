package models

type Variant struct {
	ID        string `json:"id"`
	ProductID string `json:"product_id"`
	Name      string `json:"name"`
	Details   string `json:"details"`
}
