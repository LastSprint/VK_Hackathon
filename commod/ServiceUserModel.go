package commod

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserTypeModel int

const (
	Mentor UserTypeModel = 0
	Psy    UserTypeModel = 1
)

type ServiceUser struct {
	ID          primitive.ObjectID `bson:"_id"`
	Partner     primitive.ObjectID `bson:"partner"`
	UserType    UserTypeModel      `bson:"type"`
	Name        string
	Image       string
	Description string
	Apns        string
	Messages    []MessageModel
}

type MessageModel struct {
	Recipient string `json:"recipient"`
	Text      string
	Sender    *ServiceUser `json:"sender"`
	Time      time.Time    `json:"time"`
	IsMe      bool
}
