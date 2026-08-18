package routers

import (
	"data-buku/controllers"

	"github.com/gin-gonic/gin"
)

const bookID = "/book/:bookID"

func StartServer() *gin.Engine {
	router := gin.Default()

	router.POST("/book", controllers.InsertBook)
	router.GET("/book", controllers.GetBooks)
	router.PUT(bookID, controllers.UpdateBook)
	router.DELETE(bookID, controllers.DeleteBook)

	return router
}
