package models

type Product struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	BrandName string `json:"brand_name"`
	Details   string `json:"details"`
	ImageUrl  string `json:"image_url"`
}

type ProductWithVariants struct {
	ID        string        `json:"id"`
	Name      string        `json:"name"`
	BrandName string        `json:"brand_name"`
	Details   string        `json:"details"`
	ImageUrl  string        `json:"image_url"`
	Variant   []VariantInfo `json:"variant,omitempty"`
}

type VariantInfo struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Details string `json:"details"`
}
