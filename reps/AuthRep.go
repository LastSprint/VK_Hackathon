package reps

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuthRep struct {
	cntx *DBContext
}

var userDb = "users"

type UserTypeModel int

const (
	Mentor UserTypeModel = 0
	Psy    UserTypeModel = 1
)

type ServiceUser struct {
	ID       primitive.ObjectID `bson:"_id"`
	UserType UserTypeModel
}

func InitAuthRep(cntx *DBContext) *AuthRep {
	return &AuthRep{cntx: cntx}
}

func (rep *AuthRep) GetUser(log, pass string) (*ServiceUser, error) {
	err := rep.cntx.client.Ping(rep.cntx.cntx, nil)
	fmt.Println("PING")
	fmt.Println(err)
	if err != nil {
		return nil, err
	}

	fmt.Println("COLLECTION")

	collection := rep.cntx.db.Collection(userDb)

	res := collection.FindOne(rep.cntx.cntx, bson.M{
		"$and": bson.A{
			bson.M{
				"email": log,
			},
			bson.M{
				"password": pass,
			},
		},
	})

	fmt.Println("log" + log + " " + pass)

	fmt.Println(res)

	if res.Err() != nil {
		return nil, res.Err()
	}

	fmt.Println("DECODE")

	var user ServiceUser

	fmt.Println(res)

	err = res.Decode(&user)

	if res.Err() != nil {
		return nil, err
	}

	return &user, nil
}

func (rep *AuthRep) RegToken(token string, id primitive.ObjectID) error {
	err := rep.cntx.client.Ping(rep.cntx.cntx, nil)

	if err != nil {
		return err
	}

	collection := rep.cntx.db.Collection(userDb)

	res, err := collection.UpdateOne(rep.cntx.cntx, bson.M{"_id": id}, bson.M{
		"$set": bson.M{
			"token": token,
		},
	})

	fmt.Println(res)
	fmt.Println(err)

	if err != nil {
		return err
	}

	return nil
}
