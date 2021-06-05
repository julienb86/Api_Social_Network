package controllers

import "github.com/gofiber/fiber/v2"

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
