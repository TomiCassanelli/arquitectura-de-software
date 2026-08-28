package main

// import (
// 	"context"
// 	"log"

// 	"github.com/gin-gonic/gin"

// 	"main/controllers/items"
// 	"main/database"
// 	repo "main/repositories/items"
// 	services "main/services/items"
// )

// func main() {
// 	log.Println("=== EJERCICIO 1: MongoDB Puro (Con latencia) ===")

// 	// 1. Infraestructura
// 	mongoClient := database.ConnectDB()
// 	defer mongoClient.Disconnect(context.Background())

// 	// 2. Cableado directo (Sin Caché)
// 	mongoRepo := repo.ItemsMongoDB{
// 		Client: mongoClient,
// 	}

// 	itemsService := services.ItemsService{
// 		Repo: mongoRepo, // El servicio habla directo con la BD lenta
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
