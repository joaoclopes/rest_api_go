package controller

import (
	"net/http"
	"rest_api/usecase"

	"github.com/gin-gonic/gin"
)

type itemController struct {
	itemUsecase usecase.ItemUsecase
}

func NewItemController(usecase usecase.ItemUsecase) itemController {
	return itemController{
		itemUsecase: usecase,
	}
}

func (i *itemController) GetItems(ctx *gin.Context) {
	products, err := i.itemUsecase.GetItems()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, products)
}
