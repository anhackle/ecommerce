package po

type Product struct {
	ID                int               `gorm:"primaryKey, autoIncrement, not null; unique;"`
	Name              string            `gorm:"not null"`
	Description       string            `gorm:"not null"`
	Rate              float32           `gorm:"not null"`
	ProductDetails    []ProductDetail   `gorm:"foreignKey:ProductID; references:ID"`
	ProductCategories []ProductCategory `gorm:"foreignKey:ProductID; references:ID"`
}

type ProductDetail struct {
	ID        int            `gorm:"primaryKey, autoIncrement, not null; unique;"`
	ProductID int            `gorm:"not null"`
	Color     string         `gorm:"not null"`
	Size      string         `gorm:"not null"`
	Quantity  int            `gorm:"not null"`
	Price     int            `gorm:"not null"`
	Images    []ProductImage `gorm:"foreignKey:ProductDetailID; references:ID"`
}

type ProductImage struct {
	ID              int    `gorm:"primaryKey, autoIncrement, not null; unique;"`
	ProductDetailID int    `grom:"not null"`
	URL             string `gorm:"not null"`
	AltText         string `gorm:"not null"`
}

type Category struct {
	ID                int               `gorm:"primaryKey, autoIncrement, not null; unique;"`
	Category          string            `gorm:"not null"`
	ProductCategories []ProductCategory `gorm:"foreignKey:CategoryID; references:ID"`
}

type ProductCategory struct {
	ID         int      `gorm:"primaryKey, autoIncrement, not null; unique;"`
	ProductID  int      `gorm:"not null"`
	CategoryID int      `gorm:"not null"`
	Category   Category `gorm:"foreignKey:CategoryID; references:ID"`
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

func (u *ProductCategory) TableName() string {
	return "product_category"
}
