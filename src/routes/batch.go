package routes

import (
	"backend/controller/batch"
	"backend/middlerwares"
)

func batch_routes() {

	API.Post("/create-batch", middlerwares.Auth, batch.CreateBatch)
	API.Get("/batch-list", middlerwares.Auth, batch.BatchList)
	API.Patch("/update-batch/:batchID", middlerwares.Auth, batch.UpdateBatch)

}
