package items

import (
	"encoding/json"
	"fmt"
	"main/models/items"

	"github.com/bradfitz/gomemcache/memcache"
)

// ==========================================
// EJERCICIO 3: Memcached Básico
// Muestra qué pasa si dependemos 100% de la caché
// ==========================================

type ItemsMemcachedEj3 struct {
	Client   *memcache.Client
	NextRepo ItemsRepo
}

func (repo ItemsMemcachedEj3) GetItemByID(id string) (items.ItemModel, error) {
	key := "ITEM:" + id

	// 1. Buscamos en Memcached
	cachedItem, err := repo.Client.Get(key)

	// 2. Si NO hay error -> CACHE HIT
	if err == nil {
		fmt.Println("CACHE HIT (Memcached): Devolviendo rápido")
		var item items.ItemModel
		json.Unmarshal(cachedItem.Value, &item)
		return item, nil
	}

	// 3. OJO ACÁ: Manejo de errores frágil
	// Si el error es DISTINTO a "no lo encontré" (ErrCacheMiss),
	// significa que el servidor se cayó o hay un problema de red.
	if err != memcache.ErrCacheMiss {
		fmt.Println("ERROR CRÍTICO: Memcached se cayó y rompimos la API!")
		// ESTO ROMPE EL SISTEMA: Le devolvemos el error al usuario final
		return items.ItemModel{}, err
	}

	// 4. Si solo fue un MISS normal, vamos a Mongo
	fmt.Println("CACHE MISS: Vamos a Mongo")
	item, dbErr := repo.NextRepo.GetItemByID(id)
	if dbErr != nil {
		return items.ItemModel{}, dbErr
	}

	// Guardamos en Memcached
	itemBytes, _ := json.Marshal(item)
	repo.Client.Set(&memcache.Item{
		Key:        key,
		Value:      itemBytes,
		Expiration: 60,
	})

	return item, nil
}

func (repo ItemsMemcachedEj3) CreateItem(item items.ItemModel) error {
	return repo.NextRepo.CreateItem(item)
}
