package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tracker/controller/repositories"
)

type ProductHandler struct {
	repo *repositories.ProductAPIRepository
}

func NewProductHandler(repo *repositories.ProductAPIRepository) *ProductHandler {
	return &ProductHandler{repo: repo}
}

func (h *ProductHandler) GetProducts(c *gin.Context) {
	products, err := h.repo.GetProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}

func (h *ProductHandler) GetProductByID(c *gin.Context) {
	id := c.Param("id")
	product, err := h.repo.GetProductByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if product == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	c.JSON(http.StatusOK, product)
}

func (h *ProductHandler) GetProductByName(c *gin.Context) {
	name := c.Param("name")
	products, err := h.repo.GetProductByName(name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if products == nil || len(products) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No products found with the given name"})
		return
	}
	c.JSON(http.StatusOK, products)
}
