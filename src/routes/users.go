package routes

import (
	"backend/controller/users"
	"backend/middlerwares"
)

func users_routes() {
	API.Post("/register", middlerwares.Auth, users.CreateUser)
	API.Post("/login", middlerwares.Auth, users.UserLogins)

}
