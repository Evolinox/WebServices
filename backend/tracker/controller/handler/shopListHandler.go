package handler

import "tracker/controller/repositories"

type ShopListHandler struct {
	repo *repositories.ShopListAPIRepository
}

func NewShopListHandler(repo *repositories.ShopListAPIRepository) *ShopListHandler {
	return &ShopListHandler{repo: repo}
}
