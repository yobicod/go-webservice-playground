package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Author string  `json:"author"`
	Price  float64 `json:"price"`
}

var books = []book{
	{
		ID:     "1",
		Name:   "Harry Potter",
		Author: "J.K. Rowling",
		Price:  15.9,
	},
	{
		ID:     "2",
		Name:   "One Piece",
		Author: "Oda Eiichirō",
		Price:  2.99,
	},
	{
		ID:     "3",
		Name:   "demon slayer",
		Author: "koyoharu gotouge",
		Price:  2.99,
	},
}

func getAllBooks(c *gin.Context) {
	c.JSON(http.StatusOK, books)
}

func getBookById(c *gin.Context) {
	paramID := c.Param("id")
	for _, book := range books {
		if book.ID == paramID {
			c.JSON(http.StatusOK, book)
			return
		}
	}

	// for i := 0; i <= len(books)-1; i++ {
	// 	if books[i].ID == paramID {
	// 		c.JSON(http.StatusOK, books[i])
	// 		return
	// 	}
	// }
	c.JSON(http.StatusNotFound, "Data not found")
}

func addBook(c *gin.Context) {
	var newBook book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	books = append(books, newBook)
	fmt.Println(books)
	c.JSON(http.StatusCreated, newBook)
}

func updateBook(c *gin.Context) {
	var updateBook book

	if err := c.BindJSON(&updateBook); err != nil {
		return
	}
	paramId := c.Param("id")

	// for i := 0; i <= len(books)-1; i++ {
	// 	if books[i].ID == paramId {
	// 		books[i].Name = updateBook.Name
	// 		books[i].Author = updateBook.Author
	// 		books[i].Price = updateBook.Price
	// 		c.JSON(http.StatusOK, books[i])
	// 		return
	// 	}
	// }

	for _, value := range books {
		if value.ID == paramId {
			value.Name = updateBook.Name
			value.Author = updateBook.Author
			value.Price = updateBook.Price
			c.JSON(http.StatusOK, value)
			return
		}
	}
	c.JSON(http.StatusNotFound, "Data not found")
}

func deleteBook(c *gin.Context) {
	paramId := c.Param("id")

	// for i := 0; i <= len(books)-1; i++ {
	// 	if books[i].ID == paramId {
	// 		c.JSON(http.StatusOK, "Delete success")
	// 		return
	// 	}
	// }

	for _, v := range books {
		if v.ID == paramId {
			c.JSON(http.StatusOK, "Delete success")
			return
		}
	}

	c.JSON(http.StatusNotFound, "Can't delete data(data not found)")
}

func main() {
	router := gin.Default()
	router.GET("/books", getAllBooks)
	router.GET("/book/:id", getBookById)
	router.POST("/books", addBook)
	router.PUT("/book/:id", updateBook)
	router.DELETE("/book/:id", deleteBook)
	router.Run("localhost:8080")
}
