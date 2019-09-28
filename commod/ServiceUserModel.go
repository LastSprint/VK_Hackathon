package commod

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserTypeModel int

const (
	Mentor UserTypeModel = 0
	Psy    UserTypeModel = 1
)

type ServiceUser struct {
	ID       primitive.ObjectID `bson:"_id"`
	UserType UserTypeModel
	Name     string
	Image    string
}
