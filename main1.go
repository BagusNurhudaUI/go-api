package main

// import (
// 	"fmt"

// 	"github.com/gin-gonic/gin"
// )

// type Book struct {
// 	Page int
// 	Type TypeBook
// }

// type TypeBook struct {
// 	ID    int
// 	Name  string
// 	Cover string
// }

// func (b *Book) toString() string {
// 	return fmt.Sprintf("hello this is a book %v, type: %v\n", b.Page, b.Type)
// }

// func (b *Book) addBook(book *Book) (*Book, error) {
// 	fmt.Println("entering add book ")

// 	if book.Type.Name != "book" {
// 		return nil, fmt.Errorf("Error adding book, this is not a book")
// 	}

// 	return book, nil
// }
// func main() {
// 	const version = "v1.0.0"
// 	fmt.Println("Starting the application on version \n" + version)
// 	fmt.Printf("Hello this is version %v \n\n", "1")

// 	book := &Book{
// 		Page: 1,
// 		Type: TypeBook{
// 			Name: "book",
// 		},
// 	}
// 	fmt.Printf("%v", book.toString())

// 	res, err := book.addBook(book)
// 	if err != nil {
// 		fmt.Printf("error in add book, err=%v\n", err)
// 		return
// 	}

// 	fmt.Println(res)
// 	// static.Init()

// 	// router := gin.Default()

// 	// router.GET("/ping", func(c *gin.Context) {
// 	// 	c.JSON(200, gin.H{
// 	// 		"message": "pong",
// 	// 	})
// 	// })

// 	// router.GET("/", middleware, func(c *gin.Context) {
// 	// 	c.JSON(200, gin.H{
// 	// 		"message": "Successfully connected to the server",
// 	// 		"version": version,
// 	// 	})
// 	// })

// 	// router.MaxMultipartMemory = 8 << 20 // 8 MiB
// 	// router.Static("/", "./public")
// 	// router.POST("/upload", func(c *gin.Context) {
// 	// 	name := c.PostForm("name")
// 	// 	email := c.PostForm("email")

// 	// 	// Source
// 	// 	file, err := c.FormFile("file")
// 	// 	if err != nil {
// 	// 		c.String(http.StatusBadRequest, "get form err: %s", err.Error())
// 	// 		return
// 	// 	}

// 	// 	filename := filepath.Base(file.Filename)
// 	// 	if err := c.SaveUploadedFile(file, filename); err != nil {
// 	// 		c.String(http.StatusBadRequest, "upload file err: %s", err.Error())
// 	// 		return
// 	// 	}

// 	// 	c.String(http.StatusOK, "File %s uploaded successfully with fields name=%s and email=%s.", file.Filename, name, email)
// 	// })
// 	// router.Run(":8080")
// }

// func middleware(c *gin.Context) {
// 	fmt.Println("going into the middleware")
// 	// fmt.Printf("c %v", c.Request)
// 	c.Next()
// }
