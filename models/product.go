package models

import (
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Product struct {
	// ID          string          `gorm:"size:36;not null;uniqueIndex;primaryKey;" json:"id"`
	ID uint64 `gorm:"primaryKey" json:"id"`
	// gorm.Model
	Name        string          `gorm:"size:100;not null" json:"name" binding:"required"`
	Price       decimal.Decimal `gorm:"type:decimal(16,2);" json:"price" binding:"required,numeric,min=1000"`
	Description string          `sql:"type:longtext;" json:"description,omitempty" binding:"required"`
	Image       string          `gorm:"type:text" json:"image" binding:"required"`
	CreatedAt   time.Time       `sql:"DEFAULT:'current_timestamp'" json:"created_at"`
	UpdatedAt   time.Time       `sql:"DEFAULT:ON UPDATE current_timestamp" json:"updated_at"`
	DeletedAt   gorm.DeletedAt  `gorm:"index"`
}

func (p *Product) TableName() string {
	return "products"
}
