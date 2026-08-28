package items

import (
	"main/models/items"
	repo "main/repositories/items"
	"testing"
)

func TestGetItem_Success(t *testing.T) {
	// 1. Preparamos nuestra "base de datos" de mentira
	// Inicializamos el mapa y le metemos un dato a la fuerza
	dbEnMemoria := make(map[string]items.ItemModel)
	dbEnMemoria["123"] = items.ItemModel{ID: "123", Title: "Laptop Gamer", Price: 1500.0}

	// 2. Instanciamos el Repositorio Mock en lugar del de MongoDB
	mockRepo := repo.ItemsMock{
		MockDB: dbEnMemoria,
	}

	// 3. ¡EL CAMBIO MÁGICO! Le inyectamos el Mock al Service.
	// El Service ni se entera de que esto no es MongoDB.
	service := ItemsService{
		Repo: mockRepo,
	}

	// 4. Ejecutamos la función del Service que queremos probar
	item, err := service.GetItem("123")

	// 5. Verificamos que el Service haya hecho su trabajo (Asserts)
	if err != nil {
		t.Errorf("No se esperaba un error, pero falló: %v", err)
	}

	if item.Title != "Laptop Gamer" {
		t.Errorf("Se esperaba 'Laptop Gamer', pero se obtuvo: %s", item.Title)
	}
}
