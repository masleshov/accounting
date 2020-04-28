package repository

import (
	"accounting/accounting/config"
	"accounting/accounting/data/database"
)

// DBResult represents some result of request to database - data and/or error
type DBResult struct {
	Data  interface{}
	Error error
}

// NewDBResult creates new instance of DBResult
func newDBResult(data interface{}, err error) DBResult {
	res := &DBResult{
		Data:  data,
		Error: err,
	}
	return *res
}

// Repository represents object which contains methods for working with database
type Repository struct {
	db database.Database
}

// NewRepository creates new instance of Repository
func NewRepository() *Repository {
	config := config.NewAppConfig()

	return &Repository{
		db: database.NewDatabase(config.Database.Type,
			config.Database.User,
			config.Database.Password,
			config.Database.Host,
			config.Database.Port,
			config.Database.DbName),
	}
}
