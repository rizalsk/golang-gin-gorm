package migration

import (
	"my-gin-gorm/models"
)

type Model struct {
	Model interface{}
}

func RegisterModels() []Model {
	return []Model{
		{Model: &models.User{}},
		{Model: models.Product{}},
	}
}
