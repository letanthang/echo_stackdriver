package types

import "time"

type Profile struct {
	ID         int    `json:"id" bson:"id" gorm:"column:id" groups:"internal,private,staff"`
	CustomerID int    `json:"customer_id" bson:"customer_id" gorm:"column:account_id" groups:"internal,private,staff"`
	ServiceID  int    `json:"service_id" bson:"service_id" gorm:"column:service_id" groups:"internal,private,staff"`
	Fullname   string `json:"fullname,omitempty" bson:"fullname,omitempty" groups:"internal,private,internal,staff"`
	Email      string `json:"email,omitempty" bson:"email,omitempty" groups:"private,internal,staff,public"`
	Phone      string `json:"phone,omitempty" bson:"phone,omitempty" gorm:"column:phone" groups:"private,internal,staff,public"`
	Avatar     string `json:"avatar,omitempty" bson:"avatar,omitempty" groups:"private,internal,public"`
	IsCompany  bool   `json:"is_company" bson:"is_company,omitempty" gorm:"column:is_company" groups:"private,internal,staff,public"`
}

type Profile2 struct {
	ID             int    `json:"id" bson:"id" gorm:"column:id" groups:"internal,private,staff"`
	CustomerID     int    `json:"customer_id" bson:"customer_id" gorm:"column:account_id" groups:"internal,private,staff"`
	ServiceID      int    `json:"service_id" bson:"service_id" gorm:"column:service_id" groups:"internal,private,staff"`
	Fullname       string `json:"fullname,omitempty" bson:"fullname,omitempty" groups:"internal,private,internal,staff"`
	Email          string `json:"email,omitempty" bson:"email,omitempty" groups:"private,internal,staff,public"`
	Phone          string `json:"phone,omitempty" bson:"phone,omitempty" gorm:"column:phone" groups:"private,internal,staff,public"`
	Avatar         string `json:"avatar,omitempty" bson:"avatar,omitempty" groups:"private,internal,public"`
	IsCompany      bool   `json:"is_company" bson:"is_company,omitempty" gorm:"column:is_company" groups:"private,internal,staff,public"`
	GroupAccountID int    `json:"group_account_id" bson:"group_account_id,omitempty" gorm:"column:group_account_id" groups:"public,private,internal,staff"`
}

func (Profile) TableName() string {
	return "profile"
}

type ProfileFull struct {
	CustomerID int       `json:"customer_id" bson:"customer_id" gorm:"column:account_id" groups:"internal,private,staff,public"`
	ServiceID  int       `json:"service_id" bson:"service_id" gorm:"column:service_id" groups:"internal,private,staff"`
	Phone      string    `json:"phone,omitempty" bson:"phone,omitempty" gorm:"column:phone" groups:"public,private,internal,staff"`
	Password   string    `json:"password" bson:"password" groups:"internal"`
	IsActived  bool      `json:"is_actived" bson:"is_actived" groups:"internal,staff"`
	IsDeleted  bool      `json:"is_deleted" bson:"is_deleted" groups:"internal,staff"`
	CreatedAt  time.Time `json:"created_at" bson:"created_at" groups:"private,internal,staff"`
	UpdatedAt  time.Time `json:"updated_at" bson:"updated_at" groups:"private,internal,staff"`

	Fullname       string `json:"fullname,omitempty" bson:"fullname,omitempty" groups:"public,private,internal,staff"`
	Email          string `json:"email,omitempty" bson:"email,omitempty" groups:"public,private,internal,staff"`
	Avatar         string `json:"avatar,omitempty" bson:"avatar,omitempty" groups:"public,private,internal,staff"`
	IsCompany      bool   `json:"is_company" bson:"is_company,omitempty" gorm:"column:is_company" groups:"public,private,internal,staff"`
	GroupAccountID int    `json:"group_account_id" bson:"group_account_id,omitempty" gorm:"column:group_account_id" groups:"public,private,internal,staff"`

	Contact []Contact `json:"contact" bson:"contact" groups:"public,private,internal,staff"`
	Token   string    `json:"token,omitempty" bson:"token,omitempty" groups:"public,private,internal,staff"`
}

type ProfileRequest struct {
	CustomerID *int    `json:"customer_id,omitempty"`
	FullName   *string `json:"fullname,omitempty"`
	Phone      *string `json:"phone,omitempty"`
	Email      *string `json:"email,omitempty"`
	Avatar     *string `json:"avatar,omitempty"`
	IsCompany  *bool   `json:"is_company,omitempty"`
}

type ProfileRequest2 struct {
	CustomerID     *int    `json:"customer_id,omitempty"`
	FullName       *string `json:"fullname,omitempty"`
	Phone          *string `json:"phone,omitempty"`
	Email          *string `json:"email,omitempty"`
	Avatar         *string `json:"avatar,omitempty"`
	IsCompany      *bool   `json:"is_company,omitempty"`
	GroupAccountID *int    `json:"group_account_id" gorm:"column:group_account_id"`
}
