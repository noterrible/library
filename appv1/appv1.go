package appv1

import (
	"libraryManagementSystem/appv1/model"
	"libraryManagementSystem/appv1/router"
)

func Start() {
	model.New()
	r := router.New()
	r.Run(":8080")
}
