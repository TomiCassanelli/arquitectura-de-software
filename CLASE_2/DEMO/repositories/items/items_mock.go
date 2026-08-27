package items

import (
	"errors"
	"main/models/items"
)

// ItemsMock simula ser MongoDB, pero vive solo en la RAM
type ItemsMock struct {
	MockDB map[string]items.ItemModel
}

// Implementamos el CreateItem para el Mock
func (repo ItemsMock) CreateItem(item items.ItemModel) error {
	if item.ID == "" {
		return errors.New("el ID no puede estar vacío")
	}
	// "Insertamos" el dato guardándolo en el diccionario
	repo.MockDB[item.ID] = item
	return nil
}

// Implementamos el GetItemByID para el Mock
func (repo ItemsMock) GetItemByID(id string) (items.ItemModel, error) {
	// Buscamos en el diccionario
	item, existe := repo.MockDB[id]
	if !existe {
		return items.ItemModel{}, errors.New("item no encontrado")
	}
	return item, nil
}
