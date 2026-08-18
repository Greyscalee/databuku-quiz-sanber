package routers

import (
	"data-buku/controllers"

	"github.com/gin-gonic/gin"
)

const bookID = "/book/:bookID"

func StartServer() *gin.Engine {
	router := gin.Default()

	router.POST("/bioskops", controllers.InsertBook)
	router.GET("/bioskops", controllers.GetBooks)
	router.PUT(bookID, controllers.UpdateBook)
	router.DELETE(bookID, controllers.DeleteBook)

	return router
}
