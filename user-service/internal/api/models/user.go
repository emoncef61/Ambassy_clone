package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID             primitive.ObjectID `bson:"_id,omitempy"`
	PersonalInfo   PersonalInfo       `bson:"personal_info"`
	Metadata       Metadata           `bson:"metadata"`
	ContactDetails ContactDetails     `bson:"contact_details"`
}

type PersonalInfo struct {
	FirstName   string `bson:"first_name"`
	LastName    string `bson:"last_name"`
	Email       string `bson:"email"`
	DateOfBirth string `bson:"date_of_birth"`
	Password    []byte `bson:"password" json:"-"`
}

type ContactDetails struct {
	Phone   string  `bson:"phone"`
	Address Address `bson:"address"`
}

type Metadata struct {
	CreatedAt     time.Time `bson:"created_at,omitempy"`
	LastUpdated   time.Time `bson:"last_updated,omitempy"`
	AccountStatus string    `bson:"account_status,omitempy"`
	Active        bool      `bson:"active,omitempy"`
}

type Address struct {
	Street  string `bson:"street"`
	City    string `bson:"city"`
	Country string `bson:"country"`
}
