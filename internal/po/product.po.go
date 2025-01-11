package po

type Product struct {
	ID             int             `gorm:"primaryKey, autoIncrement, not null; unique;"`
	Name           string          `gorm:"not null"`
	Description    string          `gorm:"not null"`
	Rate           float32         `gorm:"not null"`
	ProductDetails []ProductDetail `gorm:"foreignKey:ProductID; references:ID"`
	Images         []ProductImage  `gorm:"foreignKey:ProductID; references:ID"`
	Categories     []Category      `gorm:"many2many:product_category"`
}

type ProductDetail struct {
	ID        int    `gorm:"primaryKey, autoIncrement, not null; unique;"`
	ProductID int    `gorm:"not null"`
	Color     string `gorm:"not null"`
	Size      string `gorm:"not null"`
	Quantity  int    `gorm:"not null"`
	Price     int    `gorm:"not null"`
}

type ProductImage struct {
	ID        int    `gorm:"primaryKey, autoIncrement, not null; unique;"`
	ProductID int    `grom:"not null"`
	URL       string `gorm:"not null"`
	AltText   string `gorm:"not null"`
}

type Category struct {
	ID       int    `gorm:"primaryKey, autoIncrement, not null; unique;"`
	Category string `gorm:"not null"`
}

func (u *Product) TableName() string {
	return "product"
}

func (u *ProductDetail) TableName() string {
	return "product_detail"
}

func (u *ProductImage) TableName() string {
	return "product_image"
}

func (u *Category) TableName() string {
	return "category"
}
