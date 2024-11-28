package service

import (
	"database/sql"
	"remmi-cookbook/cmd/model"
)

func getRecipesForUser(userId string, db *sql.DB) ([]model.Recipe, error) {
	return make([]model.Recipe, 0), nil
}
