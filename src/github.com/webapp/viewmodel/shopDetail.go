package viewmodel

import (
	"github.com/webapp/model"
)

// ShopDetail struct
type ShopDetail struct {
	Title    string
	Active   string
	Products []Product
}

// NewShopDetail constructor
func NewShopDetail(products []model.Product) ShopDetail {
	result := ShopDetail{
		Title:    "Lemonade Stand Supply",
		Active:   "shop",
		Products: []Product{},
	}
	for _, p := range products {
		result.Products = append(result.Products, productToVM(&p))
	}
	return result
}
