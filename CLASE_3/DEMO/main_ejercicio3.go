package main

// import (
// 	"context"
// 	"log"

// 	"github.com/bradfitz/gomemcache/memcache"
// 	"github.com/gin-gonic/gin"

// 	"main/controllers/items"
// 	"main/database"
// 	repo "main/repositories/items"
// 	services "main/services/items"
// )

// func main() {
// 	log.Println("=== EJERCICIO 3: Caché Distribuida (Memcached) ===")

// 	// 1. Infraestructura
// 	mongoClient := database.ConnectDB()
// 	defer mongoClient.Disconnect(context.Background())

// 	// Apuntamos al Docker
// 	memcachedClient := memcache.New("localhost:11211")

// 	// 2. Cableado con Memcached
// 	mongoRepo := repo.ItemsMongoDB{
// 		Client: mongoClient,
// 	}

// 	// Creamos el nuevo "Escudo" distribuido
// 	memcachedRepo := repo.ItemsMemcachedEj3{
// 		Client:   memcachedClient,
// 		NextRepo: mongoRepo, // Si hay miss, va a Mongo
// 	}

// 	itemsService := services.ItemsService{
// 		Repo: memcachedRepo, // El servicio ahora habla con Memcached
// 	}

// 	itemsController := items.ItemsController{
// 		Service: itemsService,
// 	}

// 	// 3. Servidor
// 	router := gin.Default()
// 	router.GET("/items/:id", itemsController.GetByID)

// 	log.Println("🚀 Servidor corriendo en http://localhost:8080")
// 	router.Run(":8080")
// }
