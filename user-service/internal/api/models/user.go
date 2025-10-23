package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	PersonalInfo   PersonalInfo       `bson:"personal_info" json:"personal_info"`
	Metadata       Metadata           `bson:"metadata" json:"metadata"`
	ContactDetails ContactDetails     `bson:"contact_details" json:"contact_details"`
}

type PersonalInfo struct {
	FirstName   string `bson:"first_name" json:"first_name"`
	LastName    string `bson:"last_name" json:"last_name"`
	Email       string `bson:"email" json:"email"`
	DateOfBirth string `bson:"date_of_birth" json:"date_of_birth"`
	Password    string `bson:"password" json:"password"` //to fix this
}

type ContactDetails struct {
	Phone   string  `bson:"phone" json:"phone"`
	Address Address `bson:"address" json:"address"`
}

type Metadata struct {
	CreatedAt     time.Time `bson:"created_at,omitempty" json:"created_at"`
	LastUpdated   time.Time `bson:"last_updated,omitempty" json:"last_updated"`
	AccountStatus string    `bson:"account_status,omitempty" json:"account_status"`
	Active        bool      `bson:"active,omitempty" json:"active"`
}

type Address struct {
	Street  string `bson:"street" json:"street"`
	City    string `bson:"city" json:"city"`
	Country string `bson:"country" json:"country"`
}
