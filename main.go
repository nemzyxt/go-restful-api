// Author : Nemuel Wainaina

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// book struct
type book struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Author string `json:"author"`
}

// sample data
var books = []book{
	{ID: "1", Title: "Linux Security", Author: "Nemuel"},
	{ID: "2", Title: "Windows Security", Author: "Hasherezade"},
}

// fetch all books
func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

// add a new book
func addBook(c *gin.Context) {
	var newBook book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

// get book by ID
func getBookByID(c *gin.Context) {
	id := c.Param("id")

	for _, book := range books {
		if book.ID == id {
			c.IndentedJSON(http.StatusFound, book)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"msg": "Book not found"})
}
			
func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.GET("/books/:id", getBookByID)
	router.POST("/books", addBook)

	router.Run("localhost:1234")
}