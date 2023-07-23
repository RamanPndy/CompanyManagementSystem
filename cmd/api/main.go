package main

import (
	"companybuilder/initiate"
	"log"

	"github.com/ralstan-vaz/go-errors"
)

// Version ...
var Version string

// @title Company Builder API
// @version 1.1
// @description This is a sample boilerplate api server.
// @termsOfService http://swagger.io/terms/
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:80
// @BasePath /v1
func main() {
	// Initialize the app
	err := initiate.Initialize()
	if err != nil {
		newErr := errors.Get(err)
		// If an error is encountered in main its critical and the app exits
		log.Fatal(newErr)
	}
}
