package controllers

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	config "github.com/julienb86/Api_Social_Network/Config"
	models "github.com/julienb86/Api_Social_Network/Models"
	"go.mongodb.org/mongo-driver/bson"
)

type Post struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	UserId    int    `json:"userid"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	DeletedAt string `json:"deletedAt"`
}

var posts = []*Post{
	{
		Id:        1,
		Title:     "title",
		Content:   "content",
		UserId:    1,
		CreatedAt: "04/06/2021",
		UpdatedAt: "04/06/2021",
		DeletedAt: "",
	},
	{
		Id:        2,
		Title:     "second post",
		Content:   "content",
		UserId:    1,
		CreatedAt: "04/06/2021",
		UpdatedAt: "04/06/2021",
		DeletedAt: "",
	},
}

func GetAllPosts(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"posts": posts,
		},
	})
}

func CreatePost(c *fiber.Ctx) error {
	postCollection := config.MI.DB.Collection(os.Getenv("POST_COLLECTION"))

	data := new(models.Post)
	data.CreatedAt = time.Now().UTC()
	data.UpdatedAt = time.Now().UTC()
	err := c.BodyParser(&data)

	// if error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	result, err := postCollection.InsertOne(c.Context(), data)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot insert todo",
			"error":   err,
		})
	}

	// get the inserted data
	post := &models.Post{}
	query := bson.D{{Key: "_id", Value: result.InsertedID}}

	postCollection.FindOne(c.Context(), query).Decode(post)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"posts": data,
		},
	})
}
