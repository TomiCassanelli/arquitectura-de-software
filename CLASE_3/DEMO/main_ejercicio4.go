package main

import (
	"context"
	"log"
	"os"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/gin-gonic/gin"

	"main/controllers/items"
	"main/database"
	repo "main/repositories/items"
	services "main/services/items"
)

func main() {
	log.Println("=== EJERCICIO 4: Tolerancia a fallos (Degradación Elegante) ===")

	// 1. Infraestructura
	mongoClient := database.ConnectDB()
	defer mongoClient.Disconnect(context.Background())

	memcachedClient := memcache.New("localhost:11211")

	// 2. Cableado idéntico al Ejercicio 3
	mongoRepo := repo.ItemsMongoDB{
		Client: mongoClient,
	}

	memcachedRepo := repo.ItemsMemcachedEj4{
		Client:   memcachedClient,
		NextRepo: mongoRepo,
	}

	itemsService := services.ItemsService{
		Repo: memcachedRepo,
	}

	itemsController := items.ItemsController{
		Service: itemsService,
	}

	// 3. Servidor
	router := gin.Default()
	router.GET("/items/:id", itemsController.GetByID)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("🚀 Servidor corriendo en http://localhost:" + port)
	router.Run(":" + port)
}
