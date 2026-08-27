package items

import (
	"errors"
	"main/models/items"
	repo "main/repositories/items"
)

type ItemsService struct {
	// Inyectamos la interfaz, NO la base de datos real
	Repo repo.ItemsRepo
}

func (s *ItemsService) GetItem(id string) (items.ItemModel, error) {
	if id == "" {
		return items.ItemModel{}, errors.New("el ID no puede estar vacío")
	}
	return s.Repo.GetItemByID(id)
}
