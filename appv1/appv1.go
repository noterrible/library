package appv1

import (
	"libraryManagementSystem/appv1/router"
)

func Start() {
	r := router.New()
	r.Run(":8080")
}
