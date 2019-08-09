package viewmodels

type ProductGetAllVm struct {
	Id uint `json:"id"`
	Image string `json:"image"`
	ImageAlt string `json:"imageAlt"`
	Price float64 `json:"price"`
	Promotion float64 `json:"promotion"`
	ProductName string `json:"productName"`
	Description string `json:"description"`
}
