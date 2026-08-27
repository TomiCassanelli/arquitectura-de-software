package items

import "main/models/items"

type ItemsRepo interface {
	CreateItem(item items.ItemModel) error
	GetItemByID(id string) (items.ItemModel, error)
}
