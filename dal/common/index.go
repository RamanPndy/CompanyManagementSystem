package common

import (
	"companybuilder/shared"
)

func GetSchemaTable(deps *shared.Deps, tableName string) (string, string) {
	for _, dbTable := range deps.Config.Get().DB.Tables {
		if dbTable.Name == tableName {
			return dbTable.Schema, dbTable.Name
		}
	}

	return "", ""
}
