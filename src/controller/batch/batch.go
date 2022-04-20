package batch

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

var batchCollection *mongo.Collection = database.GetCollection(database.DB, "batch")

/* =================================== Create Batch =========================================*/

func CreateBatch(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	p := models.Batch{}
	if err := json.Unmarshal(c.Body(), &p); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}
	newBatch := models.Batch{
		Id:        primitive.NewObjectID(),
		BatchName: p.BatchName,
	}
	result, err := batchCollection.InsertOne(ctx, newBatch)
	fmt.Println(result)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})

	}
	return c.Status(http.StatusCreated).JSON(responses.UserResponse{Status: http.StatusCreated, Message: "success", Data: &fiber.Map{"data": "Batch create successfully!"}})

}

/*====================================================Batch List========================================================= */

func BatchList(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var class []models.Batch
	defer cancel()

	results, err := batchCollection.Find(ctx, bson.M{})

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	//reading from the db in an optimal way
	defer results.Close(ctx)
	for results.Next(ctx) {
		var singleBatchData models.Batch
		if err = results.Decode(&singleBatchData); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
		}

		class = append(class, singleBatchData)
	}

	return c.Status(http.StatusOK).JSON(
		responses.UserResponse{Status: http.StatusOK, Message: "Batch data list.", Data: &fiber.Map{"data": class}},
	)
}

/*===================================================Update Batch ========================================================================*/

func UpdateBatch(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	batchID := c.Params("batchID")
	var class models.Batch
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(batchID)

	//validate the request body
	if err := c.BodyParser(&class); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	update := bson.M{"batchname": class.BatchName, "status": class.Status}

	result, err := batchCollection.UpdateOne(ctx, bson.M{"id": objId}, bson.M{"$set": update})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	//get updated user details
	var updatedBatch models.Batch
	if result.MatchedCount == 1 {
		err := batchCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&updatedBatch)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
		}
	}

	return c.Status(http.StatusOK).JSON(responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": updatedBatch}})
}
