package model

type ProductsMainPageInput struct {
	Page     int `json:"page" binding:"required,number,gt=0"`
	PageSize int `json:"pagesize" binding:"required,number,gt=0,lt=50"`
}

type ProductsMainPageOutput struct {
	Name    string  `json:"name"`
	Rate    float32 `json:"rate"`
	Price   int     `json:"price"`
	URL     string  `json:"url"`
	AltText string  `json:"alt_text"`
}
