package mapping

import (
	"dbo/entity"
	"time"
)

type categoryFormater struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func CategoryRow(category entity.Category) categoryFormater {
	formater := categoryFormater{
		Id:          category.Id,
		Name:        category.Name,
		Description: category.Description,
		CreatedAt:   category.CreatedAt,
		UpdatedAt:   category.UpdatedAt,
	}

	return formater
}

func Category(category []entity.Category) []categoryFormater {

	if len(category) == 0 {
		return []categoryFormater{}
	}

	var categoryArray []categoryFormater
	for _, row := range category {
		accountRow := CategoryRow(row)
		categoryArray = append(categoryArray, accountRow)
	}

	return categoryArray
}
