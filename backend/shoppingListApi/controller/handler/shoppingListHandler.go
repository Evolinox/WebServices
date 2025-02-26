package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shoppingListApi/controller/repositories"
	"shoppingListApi/infrastructure/entity"
	"strings"
)

type ShoppingListHandler struct {
	repo *repositories.ShoppingListRepository
}

func NewShoppingListHandler(repo *repositories.ShoppingListRepository) *ShoppingListHandler {
	return &ShoppingListHandler{repo: repo}
}
