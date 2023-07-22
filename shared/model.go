package shared

import (
	"companybuilder/config"
	"companybuilder/pkg/clients/db"
	httpPkg "companybuilder/pkg/clients/http"
)

// VERSION keeps the version no. (commit id) for global use
var VERSION string

// Deps ... is a shared dependencies struct that contains common singletons
type Deps struct {
	Config        config.IConfig
	Database      *db.DBInstance
	HTTPRequester httpPkg.IRequest
}
