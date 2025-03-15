package app

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ryvasa/test-dna/internal/dto"
)

type HandlerInterface interface {
	GetHello(c *gin.Context)
	GetPalindrome(c *gin.Context)
	GetAll(c *gin.Context)
	GetByID(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type Handler struct {
	service ServiceInterface
}

func NewHandler(service ServiceInterface) HandlerInterface {
	return &Handler{
		service: service,
	}
}

func (h *Handler) GetHello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello Go developers"})
}

func (h *Handler) GetPalindrome(c *gin.Context) {
	text := c.Query("text")
	result := h.service.GetPalindrome(c, text)
	if result == "Not Palindrome" {
		c.JSON(http.StatusBadRequest, gin.H{"message": result})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": result})
}

func (h *Handler) GetAll(c *gin.Context) {
	result, err := h.service.GetAll(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Get all data successfully",
		"data":    result,
	})
}

func (h *Handler) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	result, err := h.service.GetById(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Get data by Id successfully",
		"data":    result,
	})
}

func (h *Handler) Create(c *gin.Context) {
	var input dto.RequestDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.service.Create(c, &input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "Created successfully",
		"data":    result,
	})
}

func (h *Handler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	var input dto.RequestDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.service.Update(c, id, &input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Updated successfully",
		"data":    result,
	})
}

func (h *Handler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	err = h.service.Delete(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Deleted successfully"})
}
