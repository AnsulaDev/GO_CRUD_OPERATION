package controllers

import (
	"api/initializers"
	"api/models"

	"github.com/gin-gonic/gin"
)

// Creating post
func PostCreate(c *gin.Context) {

	//get data off req body

	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)
	// create a post

	post := models.Post{Title: body.Body, Body: body.Title}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"post": post,
	})
}

// getting posts
func PostsIndex(c *gin.Context) {
	//get the posts
	var posts []models.Post
	initializers.DB.Find(&posts)

	// Respond with them
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

// getting post by id's
func PostShow(c *gin.Context) {
	//get single post with id
	id := c.Param("id")

	var post models.Post
	initializers.DB.First(&post, id)

	// Respond with them
	c.JSON(200, gin.H{
		"post": post,
	})
}

// updating posts
func PostsUpdate(c *gin.Context) {
	//get the id off the url
	id := c.Param("id")

	//get the data off req body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)
	//find the post were updating
	var post models.Post
	initializers.DB.First(&post, id)

	//update it
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})
	//respond with it
	c.JSON(200, gin.H{
		"post": post,
	})
}

//to delete posts

func PostDelete(c *gin.Context) {
	//get the id off the url

	id := c.Param("id")

	//delete the posts
	initializers.DB.Delete(&models.Post{}, id)

	//respond
	c.Status(200)
}
