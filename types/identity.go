package types

import "time"

type Identity struct {
	CustomerID     int       `json:"customer_id" bson:"customer_id" gorm:"column:id" groups:"public,private,internal"`
	Phone          string    `json:"phone" bson:"phone" groups:"public,private,internal"`
	Email          string    `json:"email" bson:"email" groups:"public,private,internal"`
	Password       string    `json:"password" bson:"password" groups:"internal"`
	Token          string    `json:"token,omitempty" bson:"token,omitempty" gorm:"-"  groups:"private,internal"`
	GroupAccountID int       `json:"group_account_id" bson:"group_account_id,omitempty" gorm:"column:group_account_id" groups:"public,private,internal,staff"`
	IsActived      bool      `json:"is_actived" bson:"is_actived" gorm:"column:is_actived" groups:"internal"`
	IsDeleted      bool      `json:"is_deleted" bson:"is_deleted" gorm:"column:is_deleted" groups:"internal"`
	CreatedAt      time.Time `json:"created_at" bson:"created_at" gorm:"column:created_at" groups:"private,internal"`
	UpdatedAt      time.Time `json:"updated_at" bson:"updated_at" gorm:"column:updated_at" groups:"private,internal"`
}

type LinkAccount struct {
	ID         int    `json:"id" bson:"id" gorm:"column:id" groups:"internal,private"`
	CustomerID int    `json:"customer_id" bson:"customer_id" gorm:"column:account_id" groups:"public,private,internal"`
	Platform   string `json:"platform" bson:"platform" groups:"public,private,internal"`
	Token      string `json:"token" bson:"token"  groups:"private,internal"`
	LinkId     int    `json:"link_id" bson:"link_id" gorm:"column:link_id" groups:"public,private,internal"`
	Profile    string `json:"profile" bson:"profile" gorm:"column:profile" groups:"public,private,internal"`
}

func (LinkAccount) TableName() string {
	return "link_account"
}
func (Identity) TableName() string {
	return "identity"
}

type TenantLoginBody struct {
	ClientID     int    `json:"ClientID"`
	ClientName   string `json:"ClientName"`
	UserFullName string `json:"UserFullName"`
	Token        string `json:"Token"`
	Roles        string `json:"Roles"`
	ErrorMessage string `json:"ErrorMessage"`
	SessionToken string `json:"SessionToken"`
}

type TenantRegisterBody struct {
	ClientID     int    `json:"ClientID"`
	Existed      bool   `json:"Existed"`
	Email        string `json:"Email"`
	Password     string `json:"Password"`
	ErrorMessage string `json:"ErrorMessage"`
	SessionToken string `json:"SessionToken"`
}

type TenantChangePassBody struct {
	ErrorMessage string `json:"ErrorMessage"`
	SessionToken string `json:"SessionToken"`
}

type LoginRequest struct {
	Phone    string `json:"phone" bson:"phone" validate:"required"`
	Password string `json:"password" bson:"password" validate:"required"`
}
type RegisterRequest struct {
	Phone    string `json:"phone" bson:"phone" validate:"required"`
	Password string `json:"password" bson:"password" validate:"required"`
	Email    string `json:"email" bson:"email" validate:"required"`
	Fullname string `json:"fullname" bson:"fullname" validate:"required"`
}

type ChangePassRequest struct {
	OldPassword string `json:"old_password" bson:"old_password" validate:"required"`
	NewPassword string `json:"new_password" bson:"new_password" validate:"required"`
}
type RegisterRequestV2 struct {
	Phone    string `json:"phone" bson:"phone" validate:"required,min=8"`
	Password string `json:"password" bson:"password" validate:"required,min=4"`
	Email    string `json:"email" bson:"email" validate:"required,email"`
	Fullname string `json:"fullname" bson:"fullname" validate:"required"`
	Service  string `json:"service" bson:"service" validate:"required"`
}

type LoginRequestV2 struct {
	Phone    string `json:"phone" bson:"phone" validate:"required"`
	Password string `json:"password" bson:"password" validate:"required"`
	Service  string `json:"service" bson:"service" validate:"required"`
	Agent    string `json:"agent,omitempty"`
}
