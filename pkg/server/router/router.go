// File: router.go
// Functionalities:
//	- Create and return reference to router singleton to handle all routes.
//	- Define all routes.

package router

import (
	"fmt"
	"sync"

	"github.com/voixa/backend/env"
	"github.com/voixa/backend/server/router/handlers"

	"github.com/gin-gonic/gin"
)

/* ---------------------------------------------------------- */
/*                           Router                           */
/* ---------------------------------------------------------- */
type Router struct {
	engine *gin.Engine
	routes *[]Route
}

var (
	lock   = &sync.Mutex{}
	router *Router
)

func GetRouterSingleton() *Router {
	if router == nil {
		fmt.Println("No router instance found. Creating one...")

		if env.GetConfig("GIN_MODE")[0] == "release" {
			gin.SetMode(gin.ReleaseMode)
			fmt.Println("Running in release mode.")
		} else {
			gin.SetMode(gin.DebugMode)
			fmt.Println("Running in debug mode.")
		}

		router = &Router{
			gin.Default(),
			GetRoutesSingleton(),
		}
	}
	return router
}

func (r *Router) Run() {
	for _, route := range *r.routes {
		for method, handler := range route.methods {
			r.engine.Handle(method, route.entrypoint, handler)
		}
	}

	r.engine.Run(env.GetConfig("SERVER_HOST")[0] + ":" + env.GetConfig("SERVER_PORT")[0])
}

/* ---------------------------------------------------------- */
/*                           Route                            */
/* ---------------------------------------------------------- */
type Route struct {
	// Putting pointer field in front of non-pointer field to improve performance.
	// See: https://github.com/1pkg/gopium/issues/24#issuecomment-951203325
	methods    map[string]func(c *gin.Context) // Name of HTTP method along with the handler function
	entrypoint string                          // HTTP entrypoint
}

func getEntrypoints() map[string]map[string]func(c *gin.Context) {
	return map[string]map[string]func(c *gin.Context){
		"/alive": {
			"GET": handlers.GetAlive,
		},
	}
}

var routes *[]Route

func GetRoutesSingleton() *[]Route {
	lock.Lock()
	defer lock.Unlock()

	if routes == nil {
		routes = &[]Route{}

		for entrypoint, methods := range getEntrypoints() {
			*routes = append(*routes, Route{methods, entrypoint})
		}

		return routes
	}
	return routes
}
