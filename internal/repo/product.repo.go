package repo

import (
	"context"

	"github.com/anle/codebase/global"
	"github.com/anle/codebase/internal/model"
	"github.com/anle/codebase/internal/po"
)

type IProductRepo interface {
	GetProductsForMainPage(ctx context.Context, input model.ProductsMainPageInput) (products []po.Product, err error)
}

type productRepo struct{}

func (pr *productRepo) GetProductsForMainPage(ctx context.Context, input model.ProductsMainPageInput) (products []po.Product, err error) {
	err = global.Mdb.
		Model(&po.Product{}).
		Preload("ProductDetails.Images").Preload("ProductCategories.Category").
		Limit(input.PageSize).
		Offset((input.Page - 1) * input.PageSize).
		Find(&products).Error
	if err != nil {
		return []po.Product{}, err
	}

	return products, nil
}

func NewProductRepo() IProductRepo {
	return &productRepo{}
}
