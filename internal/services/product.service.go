package service

import (
	"context"

	"github.com/anle/codebase/internal/model"
	"github.com/anle/codebase/internal/repo"
	"github.com/anle/codebase/response"
)

type IProductService interface {
	GetProductsForMainPage(ctx context.Context, input model.ProductsMainPageInput) (result int, output model.ProductsMainPageOutput, err error)
}

type productService struct {
	productRepo repo.IProductRepo
}

func (ps *productService) GetProductsForMainPage(ctx context.Context, input model.ProductsMainPageInput) (result int, output model.ProductsMainPageOutput, err error) {
	products, err := ps.productRepo.GetProductsForMainPage(ctx, input)
	if err != nil {
		return response.ErrCodeInternal, model.ProductsMainPageOutput{}, err
	}

	productInfoList := []model.Product{}

	for _, product := range products {
		productDetailList := []model.ProductDetail{}
		for _, detail := range product.ProductDetails {
			productDetailList = append(productDetailList, model.ProductDetail{
				Color:    detail.Color,
				Size:     detail.Size,
				Quantity: detail.Quantity,
				Price:    detail.Price,
				Image:    detail.Images[0].URL,
			})
		}

		productInfo := model.Product{
			Name:           product.Name,
			Description:    product.Description,
			Rate:           product.Rate,
			ProductDetails: productDetailList,
		}

		productInfoList = append(productInfoList, productInfo)
	}

	return response.ErrCodeSuccess, model.ProductsMainPageOutput{
		Products: productInfoList,
	}, nil
}

func NewProductService(productRepo repo.IProductRepo) IProductService {
	return &productService{
		productRepo: productRepo,
	}
}
