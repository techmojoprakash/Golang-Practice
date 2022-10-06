package main

// Ref : https://www.youtube.com/watch?v=bj77B59nkTQ
import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []book{
	{ID: "1", Title: "Google Inc. by Tim", Author: "Tim Cock", Quantity: 3},
	{ID: "2", Title: "Facebook Inc. by Jhan", Author: "Jhan", Quantity: 4},
	{ID: "3", Title: "Twitter Inc. by Christin", Author: "Christin", Quantity: 76},
	{ID: "4", Title: "Instagram Inc. by Balu", Author: "Balu", Quantity: 14},
}

// GET request
func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

// POST request
func createBook(c *gin.Context) {
	var newBook book
	if err := c.BindJSON(&newBook); err != nil {
		return
	}
	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)

}

// this is helper func for main func bookbyid
func getBookById(id string) (*book, error) {
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}

	}

	return nil, errors.New("book not found")
}

func booksByID(c *gin.Context) {
	id := c.Param("id")
	book, err := getBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

func main() {
	fmt.Println("Starting server...")
	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()
	router.GET("/books", getBooks)
	router.POST("/books", createBook)
	router.GET("/books/:id", booksByID)
	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	router.Run(":8081")
	// router.Run(":3000") for a hard coded port

}
