package items

import (
	"encoding/json"
	"fmt"
	"main/models/items"

	"github.com/bradfitz/gomemcache/memcache"
)

// ==========================================
// EJERCICIO 4: Memcached con Tolerancia a Fallos
// Implementa Degradación Elegante
// ==========================================

type ItemsMemcachedEj4 struct {
	Client   *memcache.Client
	NextRepo ItemsRepo
}

func (repo ItemsMemcachedEj4) GetItemByID(id string) (items.ItemModel, error) {
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

	// 3. DEGRADACIÓN ELEGANTE
	// No me importa si es un Miss o si explotó el servidor de caché.
	// Ignoramos el error, logueamos, y hacemos fallback directo a Mongo.
	fmt.Println("DEGRADACIÓN ELEGANTE: Miss o Memcached caído. Yendo a Mongo.")

	item, dbErr := repo.NextRepo.GetItemByID(id)
	if dbErr != nil {
		return items.ItemModel{}, dbErr
	}

	// 4. Intentamos guardar en Memcached (si está caído, fallará silenciosamente)
	itemBytes, _ := json.Marshal(item)
	_ = repo.Client.Set(&memcache.Item{ // Ignoramos el error del Set con el "_"
		Key:        key,
		Value:      itemBytes,
		Expiration: 60,
	})

	return item, nil
}

func (repo ItemsMemcachedEj4) CreateItem(item items.ItemModel) error {
	return repo.NextRepo.CreateItem(item)
}
