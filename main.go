package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// albume represent data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	router := gin.Default()

	// `router.GET("/albums", getAlbum)` is setting up a route in the Gin router to handle GET requests to
	// the "/albums" endpoint. When a GET request is made to this endpoint, the `getAlbum` function is
	// called. The `getAlbum` function responds with the list of all albums as JSON. This route allows
	// clients to retrieve all albums stored in the `albums` slice by accessing the "/albums" endpoint.
	router.GET("/albums", getAlbum)

	// `router.GET("/albums/:id", getAlbumByID)` is setting up a route in the Gin router to handle GET
	// requests to the "/albums/:id" endpoint. The ":id" part of the route is a parameter that allows for
	// dynamic values to be passed in the URL. When a GET request is made to this endpoint, the
	// `getAlbumByID` function is called, which locates the album in the `albums` slice whose ID matches
	// the value provided in the URL parameter, and then returns that specific album as a response.
	router.GET("/albums/:id", getAlbumByID)

	// `router.POST("/albums", postAlbums)` is setting up a route in the Gin router to handle POST
	// requests to the "/albums" endpoint. When a POST request is made to this endpoint, the `postAlbums`
	// function is called.
	router.POST("/albums", postAlbums)

	// `router.Run("localhost:8080")` is instructing the Gin router to start listening and serving HTTP
	// requests on the specified address "localhost:8080". This means that the server will be accessible
	// locally on port 8080. When this line is executed, the server will start running and be ready to
	// handle incoming requests from clients.
	router.Run("localhost:8080")
}

// getAlbums responds with the list of all albums as JSON.
func getAlbum(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
	var newAlbum album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	// `c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})` is a response being sent
	// back to the client when a specific album with the requested ID is not found.
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
