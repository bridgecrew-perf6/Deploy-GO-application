package routes

import (
	"backend/controller/moderator"
	"backend/middlerwares"
)

func Moderator_routes() {

	API.Post("/create-moderator", middlerwares.Auth, moderator.CreateModerator)
	API.Get("/moderator-list", middlerwares.Auth, moderator.ModeratorList)
	API.Patch("/update-moderator/:moderatorID", middlerwares.Auth, moderator.UpdateModerator)

}
