package routes

import (
	"backend/controller/examCategory"
	"backend/middlerwares"
)

func examcategory_routes() {

	API.Post("/create-exam-category", middlerwares.Auth, examCategory.CreateExamCategory)
	API.Get("/exam-category-list", middlerwares.Auth, examCategory.GetExamCategory)
	API.Patch("/update-exam-category/:examcategoryID", middlerwares.Auth, examCategory.UpdatExamCategory)

}
