package app

import "github.com/gin-gonic/gin"

func NewRouter(handler HandlerInterface) *gin.Engine {
	router := gin.Default()

	router.GET("/", handler.GetHello)
	router.GET("/palindrome", handler.GetPalindrome)

	router.GET("/languages", handler.GetAll)
	languages := router.Group("/language")
	{
		languages.GET("/:id", handler.GetByID)
		languages.POST("", handler.Create)
		languages.PATCH("/:id", handler.Update)
		languages.DELETE("/:id", handler.Delete)
	}

	return router
}
