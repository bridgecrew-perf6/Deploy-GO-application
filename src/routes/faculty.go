package routes

import (
	"backend/controller/faculty"
	"backend/middlerwares"
)

func faculty_routes() {

	API.Post("/create-faculty-name", middlerwares.Auth, faculty.CreateFaculty)
	API.Get("/faculty-list", middlerwares.Auth, faculty.FacultyList)
	API.Patch("/update-faculty/:facultyID", middlerwares.Auth, faculty.UpdateFaculty)

}
