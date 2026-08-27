package items

import (
	services "main/services/items"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ItemsController struct {
	Service services.ItemsService
}

func (c *ItemsController) GetByID(ctx *gin.Context) {
	id := ctx.Param("id") // Obtenemos el ID de la URL

	item, err := c.Service.GetItem(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, item)
}
