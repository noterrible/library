package appV2

import (
	"libraryManagementSystem/appV2/model"
	"libraryManagementSystem/appV2/router"
)

func Start() {
	model.New()
	r := router.New()
	r.Run("localhost:8080")
}
