package initiate

import (
	"companybuilder/apis"
	"companybuilder/config"
	"companybuilder/pkg/clients/db"
	httpPkg "companybuilder/pkg/clients/http"
	"companybuilder/shared"
)

// Initialize will initialize all the dependencies and the servers.
// Dependencies include config, external connections(grpc, http)
func Initialize() error {

	// Gets config
	conf, err := config.NewConfig()
	if err != nil {
		return err
	}

	// Initializes the DB connections
	dbInstance, err := db.NewConn(conf)
	if err != nil {
		return err
	}

	// loads all common dependencies
	dependencies := shared.Deps{
		Config:        conf,
		Database:      dbInstance,
		HTTPRequester: httpPkg.NewRequest(),
	}

	// Initializes servers
	err = apis.InitServers(&dependencies)
	if err != nil {
		return err
	}

	// Returns
	return nil
}
