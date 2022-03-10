package handlers

import (
	// "errors"
	// "fmt"
	"net/http"

	DTO "gin-boilerplate/app/dto"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func SendError(c *gin.Context) {

}

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if len(authHeader) == 0 || authHeader != "API_KEY" /* config.Config.ApiKey */ {
			c.AbortWithStatusJSON(http.StatusUnauthorized, DTO.HTTPError{Message: "authorization required header with api key `Authorization: YOUR_API_KEY`"})
		}
		c.Next()
	}
}

// GetIndex godoc
// @Summary      Show an index page
// @Description  get awailable methods
// @Tags         index
// @Accept       json
// @Produce      json
// @Success      200  {object}  DTO.Album
// @Failure 401 {object} DTO.HTTPError
// @Router       / [get]
// @Security ApiKeyAuth
func GetIndex(c *gin.Context) {

	log.Infof("Jopa")
	var newIndex = DTO.Index{
		Err: "GET / Not implemented",
		Msg: "Try: GET /albums; GET /albums/:id;  POST /album;",
	}
	c.IndentedJSON(http.StatusOK, newIndex)
}

// GetAlbums godoc
// @Summary      getAlbums responds with the list of all albums as JSON.
// @Description  getAlbums responds with the list of all albums as JSON.
// @Tags         albums
// @Accept       json
// @Produce      json
// @Success      200  {object}  DTO.Album
// @Failure 401 {object} DTO.HTTPError
// @Router       /albums [get]
// @Security ApiKeyAuth
func GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, DTO.Albums)
}

// postAlbums adds an album from JSON received in the request body.
// PostAlbums godoc
// @Summary      Add an album
// @Description  add by json account
// @Tags         albums
// @Accept       json
// @Produce      json
// @Param        account  body      DTO.Album  true  "Add account"
// @Success      201      {object}  DTO.Album
// @Failure 401 {object} DTO.HTTPError
// @Router       /albums [post]
// @Security ApiKeyAuth
func PostAlbums(c *gin.Context) {
	var newAlbum DTO.Album

	// Call BindJSON to bind the received JSON to newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	DTO.Albums = append(DTO.Albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)

}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
// GetAlbumByID godoc
// @Summary      Show an album
// @Description  get string by ID
// @Tags         albums
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Album ID"
// @Success      200  {object}  DTO.Album
// @Failure 401 {object} DTO.HTTPError
// @Failure 404 {object} DTO.HTTPError
// @Router       /albums/{id} [get]
// @Security ApiKeyAuth
func GetAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Loop through the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range DTO.Albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.AbortWithStatusJSON(http.StatusNotFound, DTO.HTTPError{Message: "not found"})
}
