package controllers

import (
	"data-buku/config"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Book struct {
	BookID   int    `json:"id"`
	bookName string `json:"boook_name"`
	author   string `json:"author"`
	publish  int    `json:"published"`
}

func InsertBook(ctx *gin.Context) {
	var newBook Book

	if err := ctx.ShouldBindJSON(&newBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	query := `INSERT INTO book_library (book_name, author, published) VALUES ($1, $2, $3) RETURNING id`
	err := config.DB.QueryRow(query, newBook.bookName, newBook.author, newBook.publish).Scan(&newBook.BookID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error_status":  "Failed to Create Data",
			"error_message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"book": newBook,
	})

}

func GetBooks(ctx *gin.Context) {
	query := `SELECT id, book_name, author, published FROM book_library ORDER BY id`
	rows, err := config.DB.Query(query)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error_status":  "Failed to Fetch Data",
			"error_message": err.Error(),
		})
		return
	}
	defer rows.Close()

	bookDatas := []Book{}
	for rows.Next() {
		var book Book
		if err := rows.Scan(&book.bookName, &book.author, &book.publish); err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error_status":  "Failed to Scan Data",
				"error_message": err.Error(),
			})
			return
		}
		bookDatas = append(bookDatas, book)
	}

	if err := rows.Err(); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error_status":  "Failed to Fetch Data",
			"error_message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"book": bookDatas,
	})
}

func UpdateBook(ctx *gin.Context) {
	bookID, err := strconv.Atoi(ctx.Param("bookID"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error_status":  "Invalid ID",
			"error_message": err.Error(),
		})
		return
	}

	var updatedBook Book
	if err := ctx.ShouldBindJSON(&updatedBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	query := `UPDATE book_library SET book_name = $1, author = $2, published = $3 WHERE id = $4`
	result, err := config.DB.Exec(query, updatedBook.bookName, updatedBook.author, updatedBook.publish, bookID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error_status":  "Failed to Update Data",
			"error_message": err.Error(),
		})
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("book with id %v not found", bookID),
		})
		return
	}

	updatedBook.BookID = bookID
	ctx.JSON(http.StatusOK, gin.H{
		"book": updatedBook,
	})
}

func DeleteBook(ctx *gin.Context) {
	bookID, err := strconv.Atoi(ctx.Param("bookID"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error_status":  "Invalid ID",
			"error_message": err.Error(),
		})
		return
	}

	query := `DELETE FROM book_library WHERE id = $1`
	result, err := config.DB.Exec(query, bookID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error_status":  "Failed to Delete Data",
			"error_message": err.Error(),
		})
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("book with id %v not found", bookID),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("book with id %v successfully deleted", bookID),
	})
}
