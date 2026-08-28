package main

// import (
// 	"context"
// 	"log"

// 	"github.com/gin-gonic/gin"
// 	"github.com/karlseguin/ccache/v3"

// 	controllers "main/controllers/items"
// 	"main/database"
// 	models "main/models/items"
// 	repo "main/repositories/items"
// 	services "main/services/items"
// )

// func main() {
// 	log.Println("=== EJERCICIO 2: Caché Local (ccache) ===")

// 	// 1. Infraestructura
// 	mongoClient := database.ConnectDB()
// 	defer mongoClient.Disconnect(context.Background())

// 	localCache := ccache.New(ccache.Configure[models.ItemModel]().MaxSize(1000))

// 	// 2. Cableado con escudo de Caché Local
// 	mongoRepo := repo.ItemsMongoDB{
// 		Client: mongoClient,
// 	}

// 	// Creamos el "Escudo"
// 	ccacheRepo := repo.ItemsCCache{
// 		NextRepo: mongoRepo, // Si hay miss, va a Mongo
// 		Cache:    localCache,
// 	}

// 	itemsService := services.ItemsService{
// 		Repo: ccacheRepo, // El servicio ahora habla con la Caché Local
// 	}

// 	itemsController := controllers.ItemsController{
// 		Service: itemsService,
// 	}

// 	// 3. Servidor
// 	router := gin.Default()
// 	router.GET("/items/:id", itemsController.GetByID)

// 	log.Println("🚀 Servidor corriendo en http://localhost:8080")
// 	router.Run(":8080")
// }
