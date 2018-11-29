package viewmodel

import (
	"fmt"

	"github.com/webapp/model"
)

// Shop struct
type Shop struct {
	Title      string
	Active     string
	Categories []Category
}

// Category struct
type Category struct {
	URL           string
	ImageURL      string
	Title         string
	Description   string
	IsOrientRight bool
}

// NewShop constructor
func NewShop(categories []model.Category) Shop {
	result := Shop{
		Title:  "Lemonade Stand Supply - Shop",
		Active: "shop",
	}
	result.Categories = make([]Category, len(categories))
	for i := 0; i < len(categories); i++ {
		vm := categoryToVM(categories[i])
		vm.IsOrientRight = i%2 == 1
		result.Categories[i] = vm
	}
	return result
}

// categoryToVM convert category to vm format
func categoryToVM(c model.Category) Category {
	return Category{
		URL:         fmt.Sprintf("/shop/%v", c.ID),
		ImageURL:    c.ImageURL,
		Title:       c.Title,
		Description: c.Description,
	}
}
