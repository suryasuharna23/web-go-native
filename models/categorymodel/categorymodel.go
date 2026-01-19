package categorymodel

import (
	"web-go-native/config"
	"web-go-native/entities"
)

func GetAll() []entities.Category {
	row, err := config.DB.Query(`SELECT * FROM categories`)

	if err != nil {
		panic(err)
	}

	defer row.Close()

	var categories []entities.Category

	for row.Next() {
		var category entities.Category
		if err := row.Scan(&category.Id, &category.Name, &category.CreatedAt, &category.UpdatedAt); err != nil {
			panic(err)
		}
		categories = append(categories, category)
	}
	return categories
}
