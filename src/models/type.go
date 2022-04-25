package models

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Error1 error

var Err1 Error1

// users type
type Users struct {
	Id        primitive.ObjectID `json:"id,omitempty"`
	FirstName string             `json:"firstName,omitempty" validate:"required"`
	LastName  string             `json:"lastName,omitempty" validate:"required"`
	Email     string             `json:"email,omitempty" validate:"required"`
	Password  string             `json:"password,omitempty" validate:"required"`
	UserType  string             `json:"userType,omitempty" validate:"required"`
	Otp       string             `json:"otp"`
	Status    bool               `json:"status"`
}

func (s Users) Stringify() []byte {
	res, err := json.Marshal(s)
	if err != nil {
		log.Fatalf("unable to convert to json data due to %s", err.Error())
		return res
	}
	return res
}

type ClassData struct {
	Id                primitive.ObjectID `json:"id,omitempty"`
	Title             string             `json:"Title"`
	Description       string             `json:"Description"`
	Thumbnail         string             `json:"Thumbnail"`
	Category          string             `json:"Category"`
	Type              string             `json:"Type"`
	FacultyId         string             `json:"FacultyId"`
	Time              string             `json:"Time"`
	Duration          string             `json:"Duration"`
	Studio            string             `json:"Studio"`
	EditorId          string             `json:"EditorId"`
	Class_Status      string             `json:"Class_Status"`
	YT_Privacy_Status string             `json:"YT_Privacy_Status"`
	YT_Channel_Name   string             `json:"YT_Channel_Name"`
	YT_Tags           string             `json:"YT_Tags"`
	YT_Playlist       string             `json:"YT_Playlist"`
	Language_pref     string             `json:"Language_pref"`
	YT_Class_Status   string             `json:"YT_Class_Status"`
	YT_Link           string             `json:"YT_Link"`
	YT_ID             string             `json:"YT_ID"`
	YT_Start_Time     string             `json:"YT_Start_Time"`
	YT_End_Time       string             `json:"YT_End_Time"`
	YT_Stream_Key     string             `json:"YT_Stream_Key"`
}

type ClassDatas struct {
	ClassDatas []ClassData `json:"ClassDatas"`
}

func (s ClassData) Stringify() []byte {
	res, err := json.Marshal(s)
	if err != nil {
		log.Fatalf("unable to convert to json data due to %s", err.Error())
		return res
	}
	return res
}

type Faculty struct {
	Id          primitive.ObjectID `json:"id,omitempty"`
	FacultyName string             `json:"facultyName"`
	Status      bool               `json:"status"`
}

type Facultys struct {
	Facultys []Faculty `json:"Facultys"`
}

func (s Faculty) Stringify() []byte {
	res, err := json.Marshal(s)
	if err != nil {
		log.Fatalf("unable to convert to json data due to %s", err.Error())
		return res
	}
	return res
}

func PanicCatcher(c *fiber.Ctx) error {
	if recover() != nil {
		fmt.Println("we got a panic")
		return c.SendStatus(500)
	}
	return nil
}

type Moderator struct {
	Id            primitive.ObjectID `json:"id,omitempty"`
	ModeratorName string             `json:"ModeratorName"`
	Status        bool               `json:"Status"`
}

type Moderators struct {
	Moderators []Moderator `json:"Moderators"`
}

func (s Moderator) Stringify() []byte {
	res, err := json.Marshal(s)
	if err != nil {
		log.Fatalf("unable to convert to json data due to %s", err.Error())
		return res
	}
	return res
}

type Batch struct {
	Id        primitive.ObjectID `json:"id,omitempty"`
	BatchName string             `json:"BatchName"`
	Status    bool               `json:"Status"`
}

type Batchs struct {
	Batchs []Batch `json:"Batchs"`
}

func (s Batch) Stringify() []byte {
	res, err := json.Marshal(s)
	if err != nil {
		log.Fatalf("unable to convert to json data due to %s", err.Error())
		return res
	}
	return res
}

type ExamCategory struct {
	Id     primitive.ObjectID `json:"id,omitempty"`
	Name   string             `json:"Name"`
	Status bool               `json:"Status"`
}

type ExamCategorys struct {
	ExamCategorys []ExamCategory `json:"ExamCategorys"`
}

func (s ExamCategory) Stringify() []byte {
	res, err := json.Marshal(s)
	if err != nil {
		log.Fatalf("unable to convert to json data due to %s", err.Error())
		return res
	}
	return res
}
