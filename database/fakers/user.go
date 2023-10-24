package fakers

import (
	"log"

	"github.com/bxcodec/faker/v4"

	"my-gin-gorm/models"
	"my-gin-gorm/utils"

	"gorm.io/gorm"
)

func UserFaker(db *gorm.DB) *models.User {
	var password string = "password"
	hashedPassword, err := utils.MakePassword(password)
	if err != nil {
		log.Panic(err)
	}

	return &models.User{
		// ID:            uuid.New().String(),
		Name:          faker.Name(),
		Email:         faker.Email(),
		Password:      hashedPassword, // password
		RememberToken: "",
		// CreatedAt:     time.Time{},
		// UpdatedAt:     time.Time{},
		// DeletedAt: gorm.DeletedAt{},
	}
}
