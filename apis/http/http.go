package http

import (
	"sync"

	"companybuilder/apis/http/auth"
	"companybuilder/apis/http/company"
	"companybuilder/apis/http/ping"
	"companybuilder/apis/middleware"
	"companybuilder/shared"

	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// StartServer starts the http server using the dependencies passed to it.
// It also initializes the routes
func StartServer(deps *shared.Deps, wg *sync.WaitGroup, fatalError chan error) error {
	defer wg.Done()

	address := deps.Config.Get().Server.HTTP.Address

	gin.SetMode(gin.DebugMode)
	router := gin.Default()

	router.Use(cors.Default())
	// Adds panic handler as a middleware
	router.Use(middleware.HandlePanic)

	// Initializes Ping routes
	ping.NewPingRoute(router)
	// Initialize all the routes
	auth.NewRoute(router, deps)
	// Initialize all the routes
	company.NewRoute(router, deps)

	// Logs server start
	log.Println("HTTP Server listening on : " + address)

	// Start the server
	err := router.Run(address)
	if err != nil {
		fatalError <- err
	}

	return nil
}
