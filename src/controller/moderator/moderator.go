package moderator

import (
	"backend/models"
	"backend/responses"
	"backend/storage/database"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var moderatorCollection *mongo.Collection = database.GetCollection(database.DB, "moderator")

/* =================================== Create Moderator =========================================*/

func CreateModerator(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	p := models.Moderator{}
	if err := json.Unmarshal(c.Body(), &p); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}
	newModerator := models.Moderator{
		Id:            primitive.NewObjectID(),
		ModeratorName: p.ModeratorName,
	}
	result, err := moderatorCollection.InsertOne(ctx, newModerator)
	fmt.Println(result)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})

	}
	return c.Status(http.StatusCreated).JSON(responses.UserResponse{Status: http.StatusCreated, Message: "success", Data: &fiber.Map{"data": "Moderator create successfully!"}})

}

/*====================================================Moderator List========================================================= */

func ModeratorList(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var class []models.Moderator
	defer cancel()

	results, err := moderatorCollection.Find(ctx, bson.M{})

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	//reading from the db in an optimal way
	defer results.Close(ctx)
	for results.Next(ctx) {
		var singleModeratorData models.Moderator
		if err = results.Decode(&singleModeratorData); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
		}

		class = append(class, singleModeratorData)
	}

	return c.Status(http.StatusOK).JSON(
		responses.UserResponse{Status: http.StatusOK, Message: "Moderator data list.", Data: &fiber.Map{"data": class}},
	)
}

/*===================================================Update Moderator ========================================================================*/

func UpdateModerator(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	moderatorID := c.Params("moderatorID")
	var class models.Moderator
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(moderatorID)

	//validate the request body
	if err := c.BodyParser(&class); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	update := bson.M{"moderatorname": class.ModeratorName, "status": class.Status}

	result, err := moderatorCollection.UpdateOne(ctx, bson.M{"id": objId}, bson.M{"$set": update})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	//get updated user details
	var updatedModerator models.Moderator
	if result.MatchedCount == 1 {
		err := moderatorCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&updatedModerator)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
		}
	}

	return c.Status(http.StatusOK).JSON(responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": updatedModerator}})
}
