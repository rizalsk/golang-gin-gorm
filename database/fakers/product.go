package fakers

import (
	"math"
	"math/rand"

	"github.com/bxcodec/faker/v4"

	"my-gin-gorm/models"

	faker2 "github.com/jaswdr/faker"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

func ProductFaker(db *gorm.DB) *models.Product {
	fake := faker2.New()
	return &models.Product{
		Name:        fake.Car().Model(),
		Price:       decimal.NewFromFloat(fakePrice()),
		Description: faker.Paragraph(),
		Image:       fake.File().FilenameWithExtension(),
	}
}

func fakePrice() float64 {
	return precision(rand.Float64()*math.Pow10(rand.Intn(8)), rand.Intn(2)+1)
}

// precision | a helper function to set precision of price
func precision(val float64, pre int) float64 {
	div := math.Pow10(pre)
	return float64(int64(val*div)) / div
}
