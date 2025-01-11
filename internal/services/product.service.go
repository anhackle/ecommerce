package service

import (
	"context"

	"github.com/anle/codebase/internal/model"
	"github.com/anle/codebase/internal/repo"
	"github.com/anle/codebase/response"
)

type IProductService interface {
	GetProductsForMainPage(ctx context.Context, input model.ProductsMainPageInput) (result int, output []model.ProductsMainPageOutput, err error)
}

type productService struct {
	productRepo repo.IProductRepo
}

func (ps *productService) GetProductsForMainPage(ctx context.Context, input model.ProductsMainPageInput) (result int, output []model.ProductsMainPageOutput, err error) {
	products, err := ps.productRepo.GetProductsForMainPage(ctx, input)
	if err != nil {
		return response.ErrCodeInternal, []model.ProductsMainPageOutput{}, err
	}

	for _, product := range products {
		output = append(output, model.ProductsMainPageOutput{
			Name:  product.Name,
			Rate:  product.Rate,
			Price: product.ProductDetails[1].Price,
		})
	}

}

func NewProductService(productRepo repo.IProductRepo) IProductService {
	return &productService{
		productRepo: productRepo,
	}
}
