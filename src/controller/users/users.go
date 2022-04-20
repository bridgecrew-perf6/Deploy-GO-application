package users

import (
	"backend/models"
	"backend/responses"
	"backend/storage/database"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var userCollection *mongo.Collection = database.GetCollection(database.DB, "users")

/* =================================== Create User =========================================*/

func CreateUser(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var user models.Users
	p := models.Users{}
	if err := json.Unmarshal(c.Body(), &p); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}
	err := userCollection.FindOne(ctx, bson.M{"email": p.Email}).Decode(&user)
	if err == nil {
		return c.Status(401).JSON(responses.UserResponse{Status: 401, Message: "error", Data: &fiber.Map{"data": "Email already exists."}})
	}
	password := []byte(p.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	var strToConvert string

	strToConvert = string(hashedPassword)
	newUser := models.Users{
		Id:        primitive.NewObjectID(),
		FirstName: p.FirstName,
		LastName:  p.LastName,
		Email:     p.Email,
		Password:  strToConvert,
		UserType:  p.UserType,
	}

	result, err := userCollection.InsertOne(ctx, newUser)
	fmt.Println(result)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})

	}
	return c.Status(http.StatusCreated).JSON(responses.UserResponse{Status: http.StatusCreated, Message: "success", Data: &fiber.Map{"data": "User Registered Successfully!"}})

}

/* =================================== Users Login =========================================*/
func UserLogins(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var user models.Users
	defer cancel()

	p := models.Users{}
	if err := json.Unmarshal(c.Body(), &p); err != nil {
		log.Println(err)
	}
	err := userCollection.FindOne(ctx, bson.M{"email": p.Email}).Decode(&user)
	if err != nil {
		return c.Status(401).JSON(responses.UserResponse{Status: 401, Message: "error", Data: &fiber.Map{"data": "Email doesn't exist."}})
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(p.Password))
	if err != nil {
		return c.Status(401).JSON(responses.UserResponse{Status: 401, Message: "error", Data: &fiber.Map{"data": "Password doesn't matched"}})
	}

	return c.Status(http.StatusOK).JSON(responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": user}})
}
