package examCategory

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

var examCategoryCollection *mongo.Collection = database.GetCollection(database.DB, "examcategory")

/* =================================== Create Exam Category =========================================*/

func CreateExamCategory(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	p := models.ExamCategory{}
	if err := json.Unmarshal(c.Body(), &p); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}
	newExamCategory := models.ExamCategory{
		Id:   primitive.NewObjectID(),
		Name: p.Name,
	}
	result, err := examCategoryCollection.InsertOne(ctx, newExamCategory)
	fmt.Println(result)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})

	}
	return c.Status(http.StatusCreated).JSON(responses.UserResponse{Status: http.StatusCreated, Message: "success", Data: &fiber.Map{"data": "ExamCategory create successfully!"}})

}

/*====================================================Exam Category List========================================================= */

func GetExamCategory(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var class []models.ExamCategory
	defer cancel()

	results, err := examCategoryCollection.Find(ctx, bson.M{})

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	//reading from the db in an optimal way
	defer results.Close(ctx)
	for results.Next(ctx) {
		var singleExamCategoryData models.ExamCategory
		if err = results.Decode(&singleExamCategoryData); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
		}

		class = append(class, singleExamCategoryData)
	}

	return c.Status(http.StatusOK).JSON(
		responses.UserResponse{Status: http.StatusOK, Message: "ExamCategory data list.", Data: &fiber.Map{"data": class}},
	)
}

/*===================================================Update Exam Category ========================================================================*/

func UpdatExamCategory(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	examcategoryID := c.Params("examcategoryID")
	var class models.ExamCategory
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(examcategoryID)

	//validate the request body
	if err := c.BodyParser(&class); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	update := bson.M{"name": class.Name, "status": class.Status}

	result, err := examCategoryCollection.UpdateOne(ctx, bson.M{"id": objId}, bson.M{"$set": update})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	//get updated user details
	var updatedExamCategory models.ExamCategory
	if result.MatchedCount == 1 {
		err := examCategoryCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&updatedExamCategory)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
		}
	}

	return c.Status(http.StatusOK).JSON(responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": updatedExamCategory}})
}
