package mapping

import (
	"dbo/entity"
	"time"
)

type productsFormater struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Stock       int       `json:"stock"`
	Price       int       `json:"price"`
	CategoryId  int       `json:"categoryId"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func ProductsRow(products entity.Products) productsFormater {
	formater := productsFormater{
		Id:          products.Id,
		Name:        products.Name,
		Description: products.Description,
		Stock:       products.Stock,
		Price:       products.Price,
		CategoryId:  products.CategoryId,
		CreatedAt:   products.CreatedAt,
		UpdatedAt:   products.UpdatedAt,
	}

	return formater
}

func Products(products []entity.Products) []productsFormater {

	if len(products) == 0 {
		return []productsFormater{}
	}

	var productsArray []productsFormater
	for _, row := range products {
		productsRow := ProductsRow(row)
		productsArray = append(productsArray, productsRow)
	}

	return productsArray
}
