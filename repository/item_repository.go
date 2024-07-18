package repository

import (
	"database/sql"
	"fmt"
	"rest_api/model"
)

type ItemRepository struct {
	connection *sql.DB
}

func NewItemRepository(connection *sql.DB) ItemRepository {
	return ItemRepository{
		connection: connection,
	}
}

func (ir *ItemRepository) GetItems() ([]model.Item, error) {
	query := "SELECT id, name, description, price FROM item"
	rows, err := ir.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []model.Item{}, err
	}

	var itemList []model.Item
	var itemObj model.Item

	for rows.Next() {
		err = rows.Scan(
			&itemObj.ID,
			&itemObj.Name,
			&itemObj.Description,
			&itemObj.Price)

		if err != nil {
			fmt.Println(err)
			return []model.Item{}, err
		}

		itemList = append(itemList, itemObj)
	}

	rows.Close()

	return itemList, nil
}
