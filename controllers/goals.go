package controllers

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	"github.com/gofiber/fiber/v2"
	"github.com/pavangbhat/gotest/models"
	"google.golang.org/api/iterator"
)

var (
	projectId = "golang-test-cdf69"
	collectionName="goals"
)

func GetGoals(c *fiber.Ctx) error{
	ctx := context.Background()
	client, err := firestore.NewClient(ctx,  projectId)

	if err != nil {
		log.Fatalf("Failed to create a Firestore Client: %v", err)
		c.Status(400)
		return c.JSON(fiber.Map{
			"message":"Failed to create a Firestore Client:",
		})
	}

	defer client.Close()

	var goals []models.Goals

	it := client.Collection(collectionName).Documents(ctx)
	for {
		doc, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate the list of posts: %v", err)
			break
		}
		goal := models.Goals{
			Id:    doc.Data()["id"].(int64),
			Title: doc.Data()["title"].(string),
			Status:  doc.Data()["status"].(string),
		}

		goals = append(goals, goal)
	}

	if err != nil {
		log.Fatalf("Failed geting goals: %v", err)
	}

	return c.JSON(goals)
}

func CreateGoal(c *fiber.Ctx) error{
	c.Accepts("application/json")
	ctx := context.Background()
	
	newGoal := new(models.Goals)

	if err := c.BodyParser(newGoal); err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message":"Invaild data or wrong format",
		  })
	}

	client, err := firestore.NewClient(ctx,  projectId)
	if err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message":"Failed to connect client",
		  })
	}

	defer client.Close()


	_, _, err = client.Collection(collectionName).Add(ctx, map[string]interface{}{
		"id":    newGoal.Id,
		"title": newGoal.Title,
		"status":  newGoal.Status,
	})

	if err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message":"Error writing data",
		  })
	}
	return c.JSON(fiber.Map{
		"message":"Sucessfull add a goal",
	  })
}

func UpdateGoal(c *fiber.Ctx) error{
	return c.JSON(fiber.Map{
		"message":"In progress",
	  })
}

func DeleteGoal(c *fiber.Ctx) error{
	return c.JSON(fiber.Map{
		"message":"In progress",
	  })
}