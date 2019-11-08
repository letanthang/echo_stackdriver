package types

import (
	"time"

	"github.com/lib/pq"
)

type Contact struct {
	ID         int             `json:"id" bson:"id" groups:"public,private,internal" `
	CustomerID int             `json:"account_id" bson:"account_id" gorm:"column:account_id" groups:"public,private,internal" `
	Type       int             `json:"type" bson:"type" groups:"public,private,internal" `
	Fullname   string          `json:"fullname" bson:"fullname" groups:"public,private,internal" `
	Phone      string          `json:"phone" bson:"phone" groups:"public,private,internal" `
	Email      string          `json:"email,omitempty" bson:"email,omitempty" groups:"public,private,internal" `
	Address    string          `json:"address" bson:"address" groups:"public,private,internal" `
	Location   pq.Float64Array `json:"location,omitempty" bson:"location,omitempty"  groups:"public,private,internal"`
	RegionID   int             `json:"region_id,omitempty" bson:"region_id,omitempty" gorm:"column:region_id" groups:"public,private,internal"`
	District   string          `json:"district,omitempty" bson:"district,omitempty" groups:"public,private,internal"`
	Province   string          `json:"province,omitempty" bson:"province,omitempty" groups:"public,private,internal"`
	CreatedAt  time.Time       `json:"created_at" bson:"created_at" gorm:"column:created_at" groups:"private,internal"`
	UpdatedAt  time.Time       `json:"updated_at" bson:"updated_at" gorm:"column:updated_at" groups:"private,internal"`
}

func (Contact) TableName() string {
	return "contact"
}

type ContactRequest struct {
	CustomerID int       `json:"customer_id,omitempty"`
	Type       int       `json:"type" validate:"required"`
	Fullname   string    `json:"fullname" validate:"required"`
	Phone      string    `json:"phone" validate:"required"`
	Email      string    `json:"email" validate:"required,email"`
	Address    string    `json:"address" validate:"required"`
	Location   []float64 `json:"location" validate:"required"`
	RegionID   int       `json:"region_id" validate:"required"`
	District   string    `json:"district" validate:"required"`
	Province   string    `json:"province" validate:"required"`
}

type ContactUpdateRequest struct {
	ID         int        `json:"id,omitempty" bson:"id,omitempty"`
	CustomerID int        `json:"customer_id,omitempty" bson:"customer_id,omitempty"`
	Type       int        `json:"type" bson:"type" validate:"required" `
	Fullname   *string    `json:"fullname,omitempty" bson:"fullname"`
	Phone      *string    `json:"phone,omitempty" bson:"phone"`
	Email      *string    `json:"email,omitempty" bson:"email"`
	Address    *string    `json:"address,omitempty" bson:"address"`
	Location   *[]float64 `json:"location,omitempty" bson:"location"`
	RegionID   *int       `json:"region_id,omitempty" bson:"region_id"`
	District   *string    `json:"district,omitempty" bson:"district,omitempty"`
	Province   *string    `json:"province,omitempty" bson:"province,omitempty"`
	UpdatedAt  time.Time  `json:"updated_at,omitempty" bson:"updated_at"`
}

type ContactRemoveRequest struct {
	ID         int `json:"id,omitempty" bson:"id,omitempty"`
	CustomerID int `json:"customer_id,omitempty" bson:"customer_id,omitempty"`
	Type       int `json:"type" bson:"type" validate:"required" `
}
