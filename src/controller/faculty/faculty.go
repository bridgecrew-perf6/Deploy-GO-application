package faculty

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

var facultyCollection *mongo.Collection = database.GetCollection(database.DB, "faculty")

/* =================================== Create Faculty =========================================*/

func CreateFaculty(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	p := models.Faculty{}
	if err := json.Unmarshal(c.Body(), &p); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}
	newFaculty := models.Faculty{
		Id:          primitive.NewObjectID(),
		FacultyName: p.FacultyName,
	}
	result, err := facultyCollection.InsertOne(ctx, newFaculty)
	fmt.Println(result)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})

	}
	return c.Status(http.StatusCreated).JSON(responses.UserResponse{Status: http.StatusCreated, Message: "success", Data: &fiber.Map{"data": "Faculty create successfully!"}})

}

/*====================================================Faculty List========================================================= */

func FacultyList(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var class []models.Faculty
	defer cancel()

	results, err := facultyCollection.Find(ctx, bson.M{})

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	//reading from the db in an optimal way
	defer results.Close(ctx)
	for results.Next(ctx) {
		var singleFacultyData models.Faculty
		if err = results.Decode(&singleFacultyData); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
		}

		class = append(class, singleFacultyData)
	}

	return c.Status(http.StatusOK).JSON(
		responses.UserResponse{Status: http.StatusOK, Message: "Faculty data list.", Data: &fiber.Map{"data": class}},
	)
}

/*===================================================Update Faculty ========================================================================*/

func UpdateFaculty(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	facultyID := c.Params("facultyID")
	var class models.Faculty
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(facultyID)

	//validate the request body
	if err := c.BodyParser(&class); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	update := bson.M{"facultyname": class.FacultyName, "status": class.Status}

	result, err := facultyCollection.UpdateOne(ctx, bson.M{"id": objId}, bson.M{"$set": update})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	//get updated user details
	var updatedFaculty models.Faculty
	if result.MatchedCount == 1 {
		err := facultyCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&updatedFaculty)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
		}
	}

	return c.Status(http.StatusOK).JSON(responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": updatedFaculty}})
}
