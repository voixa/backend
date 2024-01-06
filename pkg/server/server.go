package server

import (
	"github.com/voixa/backend/server/router"
)

func Start() {
	router := router.GetRouterSingleton()
	router.Run()
}
