package model

type ProductsMainPageInput struct {
	Page     int `form:"page" binding:"required,number,gt=0"`
	PageSize int `form:"pagesize" binding:"required,number,gt=0,lt=9"`
}

type ProductDetail struct {
	Color    string `json:"color"`
	Size     string `json:"size"`
	Quantity int    `json:"quantity"`
	Price    int    `json:"price"`
	Image    string `json:"image"`
}
type Product struct {
	Name           string          `json:"name"`
	Description    string          `json:"description"`
	ProductDetails []ProductDetail `json:"productDetail"`
	Rate           float32         `json:"rate"`
}
type ProductsMainPageOutput struct {
	Products []Product `json:"products"`
}
