package appV1

import (
	"libraryManagementSystem/appV1/model"
	"libraryManagementSystem/appV1/router"
)

func Start() {
	model.New()
	r := router.New()
	r.Run(":8080")
}
