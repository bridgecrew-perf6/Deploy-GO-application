package routes

import (
	"backend/middlerwares"
	"backend/pkg/youtube"
)

func youtube_routes() {

	API.Post("/create-stream/:title", middlerwares.Auth, youtube.CreateStreamYoutube)
	API.Post("/bind-stream", middlerwares.Auth, youtube.BindStreamYoutube)
	API.Post("/transition-stream", middlerwares.Auth, youtube.TransitionStreamYoutube)

}
