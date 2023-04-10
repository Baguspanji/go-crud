package controllers

import (
	"github.com/baguspanji/go-crud/initializers"
	"github.com/baguspanji/go-crud/models"
	"github.com/gin-gonic/gin"
)

func GetPosts(c *gin.Context) {
	posts := []models.Post{}

	result := initializers.DB.Find(&posts)

	if result.Error != nil {
		c.JSON(500, gin.H{
			"message": "error",
			"error":   result.Error,
		})
		return
	}

	postResponses := []models.PostResponse{}

	for _, post := range posts {
		postResponse := models.PostResponse{
			ID:        post.ID,
			Title:     post.Title,
			Content:   post.Content,
			CreatedAt: post.CreatedAt,
			UpdatedAt: post.UpdatedAt,
			DeletedAt: post.DeletedAt,
		}

		postResponses = append(postResponses, postResponse)
	}

	c.JSON(200, gin.H{
		"message": "success",
		"data":    postResponses,
	})
}

func CreatePost(c *gin.Context) {
	body := models.PostRequest{}

	err := c.BindJSON(&body)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "error",
			"errors":  err.Error(),
		})
		return
	}

	post := models.Post{Title: body.Title, Content: body.Content}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.JSON(500, gin.H{
			"message": "error",
			"error":   result.Error,
		})
		return
	}

	c.JSON(201, gin.H{
		"message": "data created",
		"data":    post,
	})
}

func GetPost(c *gin.Context) {
	id := c.Params.ByName("id")

	post := models.Post{}

	result := initializers.DB.First(&post, id)

	if result.Error != nil {
		c.JSON(404, gin.H{
			"message": "error",
			"error":   result.Error,
		})
		return
	}

	postResponse := models.PostResponse{
		ID:        post.ID,
		Title:     post.Title,
		Content:   post.Content,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
		DeletedAt: post.DeletedAt,
	}

	c.JSON(200, gin.H{
		"message": "success",
		"data":    postResponse,
	})
}

func UpdatePost(c *gin.Context) {
	id := c.Params.ByName("id")

	post := models.Post{}

	result := initializers.DB.First(&post, id)

	if result.Error != nil {
		c.JSON(404, gin.H{
			"message": "error",
			"error":   result.Error,
		})
		return
	}

	body := models.PostRequest{}

	err := c.BindJSON(&body)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "error",
			"errors":  err.Error(),
		})
		return
	}

	result = initializers.DB.Model(&post).Updates(models.Post{Title: body.Title, Content: body.Content})

	if result.Error != nil {
		c.JSON(500, gin.H{
			"message": "error",
			"error":   result.Error,
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "data updated",
		"data":    post,
	})
}

func DeletePost(c *gin.Context) {
	id := c.Params.ByName("id")

	post := models.Post{}

	result := initializers.DB.First(&post, id)

	if result.Error != nil {
		c.JSON(404, gin.H{
			"message": "error",
			"error":   result.Error,
		})
		return
	}

	result = initializers.DB.Delete(&post)

	if result.Error != nil {
		c.JSON(500, gin.H{
			"message": "error",
			"error":   result.Error,
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "data deleted",
	})
}
