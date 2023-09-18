package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"study/internal/sources"
	author "study/internal/sources/author"
)

var repo = sources.GetRepo()

func CreateAuthor(c *gin.Context) {
	auth := author.Author{}
	if err := c.ShouldBind(&auth); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err := repo.Author.Create(c, &auth)
	if err != nil {
		log.Println(err)
	}

	c.JSON(http.StatusCreated, auth)
}

// GetAuthors returns all authors
func GetAuthors(c *gin.Context) {

	authors, err := repo.Author.FindAll(c)
	if err != nil {
		log.Println(err)
	}

	c.JSON(http.StatusOK, authors)
}

// GetAuthorByID returns author by ID
func GetAuthorByID(c *gin.Context) {

	auth := author.Author{}
	if err := c.BindUri(&auth); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	authorFound, err := repo.Author.FindOne(c, auth.ID)
	if err != nil {
		log.Println(err)
	}

	c.JSON(http.StatusOK, authorFound)
}

func DeleteAuthor(c *gin.Context) {
	auth := author.Author{}
	if err := c.BindUri(&auth); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err := repo.Author.Delete(c, auth.ID)
	if err != nil {
		log.Println(err)
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}

func UpdateAuthor(c *gin.Context) {
	auth := author.Author{}

	if err := c.BindUri(&auth); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := c.ShouldBind(&auth); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err := repo.Author.Update(c, auth)
	if err != nil {
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, auth)
}
