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

// GetProducts godoc
// @Summary      List all products
// @Description  Retrieves a list of all available products.
// @Tags         products
// @Accept       json
// @Produce      json
// @Success      200  {array}   dto.ProductDTO
// @Failure      500  {object}  httputil.HTTPError
// @Router       /products/ [get]

func (h *ProductHandler) GetProducts(c *gin.Context) {
	products, err := h.repo.GetProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}

// GetProductByID godoc
// @Summary      Get a product by ID
// @Description  Retrieves details of a specific product using its ID.
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Product ID"
// @Success      200  {object}  dto.ProductDTO
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /products/{id} [get]

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

// GetProductByName godoc
// @Summary      Search products by name
// @Description  Retrieves a list of products matching a given name.
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        name  path      string  true  "Product Name"
// @Success      200   {array}   dto.ProductDTO
// @Failure      404   {object}  httputil.HTTPError
// @Failure      500   {object}  httputil.HTTPError
// @Router       /products/name/{name} [get]

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
