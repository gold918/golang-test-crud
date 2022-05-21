package config

import "github.com/upper/db/v4/adapter/postgresql"

func GetDatabaseSetting() *postgresql.ConnectionURL {
	return &postgresql.ConnectionURL{
		Database: `db_golang`,
		Host:     `localhost:5432`,
		User:     `postgres`,
		Password: `postgres`,
	}
}
