package items

import (
	"fmt"
	"main/models/items"
	"time"

	"github.com/karlseguin/ccache/v3"
)

// ==========================================
// EJERCICIO 2: Caché Local con ccache
// ==========================================

type ItemsCCache struct {
	// NextRepo es el repositorio que sigue en la cadena (Mongo)
	NextRepo ItemsRepo
	Cache    *ccache.Cache[items.ItemModel]
}

func (repo ItemsCCache) GetItemByID(id string) (items.ItemModel, error) {
	key := "ITEM:" + id

	// 1. Buscamos en la caché local (HIT)
	cached := repo.Cache.Get(key)
	if cached != nil && !cached.Expired() {
		fmt.Println("⚡ CACHE HIT (Local): Devolviendo al instante")
		return cached.Value(), nil
	}

	// 2. Si no está (MISS), delegamos al siguiente repositorio (Mongo)
	fmt.Println("🐢 CACHE MISS (Local): Yendo a la base de datos...")
	item, err := repo.NextRepo.GetItemByID(id)
	if err != nil {
		return items.ItemModel{}, err
	}

	// 3. Guardamos en caché para la próxima
	repo.Cache.Set(key, item, 5*time.Minute)
	return item, nil
}

func (repo ItemsCCache) CreateItem(item items.ItemModel) error {
	return repo.NextRepo.CreateItem(item)
}
