package routes

import (
	"backend/controller/classData"
	"backend/middlerwares"
)

func class_routes() {
	API.Post("/create-class", middlerwares.Auth, classData.CreateClass)
	API.Get("/class-list", middlerwares.Auth, classData.GetAllClass)
	API.Get("/class-data/:classId", middlerwares.Auth, classData.GetClass)
	API.Patch("/update-class/:classId", middlerwares.Auth, classData.UpdateClassStatus)

}
