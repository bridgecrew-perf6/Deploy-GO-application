package classData

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

var classCollection *mongo.Collection = database.GetCollection(database.DB, "classData")

/* =================================== Create Class =========================================*/

func CreateClass(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	p := models.ClassData{}
	if err := json.Unmarshal(c.Body(), &p); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}
	newClass := models.ClassData{
		Id:                primitive.NewObjectID(),
		Title:             p.Title,
		Description:       p.Description,
		Thumbnail:         p.Thumbnail,
		Category:          p.Category,
		Type:              p.Type,
		FacultyId:         p.FacultyId,
		Time:              p.Time,
		Duration:          p.Duration,
		Studio:            p.Studio,
		EditorId:          p.EditorId,
		Class_Status:      p.Class_Status,
		YT_Privacy_Status: p.YT_Privacy_Status,
		YT_Channel_Name:   p.YT_Channel_Name,
		YT_Tags:           p.YT_Tags,
		YT_Playlist:       p.YT_Playlist,
		Language_pref:     p.Language_pref,
		YT_Class_Status:   p.YT_Class_Status,
		YT_Link:           p.YT_Link,
		YT_ID:             p.YT_ID,
		YT_Start_Time:     p.YT_Start_Time,
		YT_End_Time:       p.YT_End_Time,
		YT_Stream_Key:     p.YT_Stream_Key,
	}
	result, err := classCollection.InsertOne(ctx, newClass)
	fmt.Println(result)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})

	}
	return c.Status(http.StatusCreated).JSON(responses.UserResponse{Status: http.StatusCreated, Message: "success", Data: &fiber.Map{"data": "Class create Successfully!"}})

}

/*====================================================Class List========================================================= */

func GetAllClass(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var class []models.ClassData
	defer cancel()

	results, err := classCollection.Find(ctx, bson.M{})

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	//reading from the db in an optimal way
	defer results.Close(ctx)
	for results.Next(ctx) {
		var singleClassData models.ClassData
		if err = results.Decode(&singleClassData); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
		}

		class = append(class, singleClassData)
	}

	return c.Status(http.StatusOK).JSON(
		responses.UserResponse{Status: http.StatusOK, Message: "Class data list.", Data: &fiber.Map{"data": class}},
	)
}

/*============================================ Class Data==================================================================================*/

func GetClass(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	classId := c.Params("classId")
	var class models.ClassData
	defer cancel()
	objId, _ := primitive.ObjectIDFromHex(classId)
	err := classCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&class)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusOK).JSON(responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": class}})
}

/*===================================================Update Class ========================================================================*/

func UpdateClassStatus(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	classId := c.Params("classId")
	var class models.ClassData
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(classId)

	//validate the request body
	if err := c.BodyParser(&class); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	update := bson.M{"class_status": class.Class_Status}

	result, err := classCollection.UpdateOne(ctx, bson.M{"id": objId}, bson.M{"$set": update})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	//get updated user details
	var updatedClass models.ClassData
	if result.MatchedCount == 1 {
		err := classCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&updatedClass)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
		}
	}

	return c.Status(http.StatusOK).JSON(responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": updatedClass}})
}
