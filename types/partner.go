package types

import "time"

type GetTokenRequest struct {
	Key string `json:"api_key" validate:"required"`
}

type PartnerApiKey struct {
	CustomerID  int       `json:"customer_id" bson:"customer_id" gorm:"column:account_id" groups:"private,internal"`
	PartnerID   int       `json:"partner_id" bson:"partner_id" gorm:"column:partner_id" groups:"private,internal"`
	PartnerName string    `json:"partner_name,omitempty" bson:"partner_name,omitempty" gorm:"column:partner_name" groups:"private,internal"`
	Key         string    `json:"key,omitempty" bson:"key,omitempty" groups:"private,internal"`
	IsActived   bool      `json:"is_actived" bson:"is_actived" gorm:"column:is_actived" groups:"internal"`
	IsDeleted   bool      `json:"is_deleted" bson:"is_deleted" gorm:"column:is_deleted" groups:"internal"`
	CreatedAt   time.Time `json:"created_at" bson:"created_at" groups:"private,internal"`
	UpdatedAt   time.Time `json:"updated_at" bson:"updated_at" groups:"private,internal"`
}

type FullPartnerApiKey struct {
	CustomerID  int       `json:"customer_id" bson:"customer_id" gorm:"column:account_id" groups:"private,internal"`
	GroupID     int       `json:"group_id"  gorm:"column:group_account_id" groups:"private,internal"`
	PartnerID   int       `json:"partner_id" bson:"partner_id" gorm:"column:partner_id" groups:"private,internal"`
	PartnerName string    `json:"partner_name,omitempty" bson:"partner_name,omitempty" gorm:"column:partner_name" groups:"private,internal"`
	Key         string    `json:"key,omitempty" bson:"key,omitempty" groups:"private,internal"`
	IsActived   bool      `json:"is_actived" bson:"is_actived" gorm:"column:is_actived" groups:"internal"`
	IsDeleted   bool      `json:"is_deleted" bson:"is_deleted" gorm:"column:is_deleted" groups:"internal"`
	CreatedAt   time.Time `json:"created_at" bson:"created_at" groups:"private,internal"`
	UpdatedAt   time.Time `json:"updated_at" bson:"updated_at" groups:"private,internal"`
}

func (PartnerApiKey) TableName() string {
	return "api_key"
}

type PartnerApiKeyRequest struct {
	CustomerID  int    `json:"customer_id" bson:"customer_id" gorm:"column:account_id" groups:"private,internal"`
	PartnerID   int    `json:"partner_id" bson:"partner_id" groups:"private,internal"`
	PartnerName string `json:"partner_name,omitempty" bson:"partner_name,omitempty" groups:"private,internal"`
}
