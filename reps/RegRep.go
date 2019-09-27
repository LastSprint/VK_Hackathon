package reps

import (
	reg "suncity/reg/models"

	"go.mongodb.org/mongo-driver/bson"
)

const dbName = "users"

// RegRep репозиторий регистрации
type RegRep struct {
	cntx *DBContext
}

// InitRegRep иницаллизирует репозиторий
func InitRegRep(cntx *DBContext) *RegRep {
	return &RegRep{cntx: cntx}
}

// PostForm регистрирует форму
func (rep *RegRep) PostForm(form reg.FormModel) error {

	err := rep.cntx.client.Ping(rep.cntx.cntx, nil)

	if err != nil {
		return err
	}

	collection := rep.cntx.db.Collection(dbName)

	_, err = collection.InsertOne(rep.cntx.cntx, form)

	return err
}

func (rep *RegRep) CheckFormStatus(value string) (*reg.CheckStatusResponseModel, error) {

	err := rep.cntx.client.Ping(rep.cntx.cntx, nil)

	if err != nil {
		return nil, err
	}

	collection := rep.cntx.db.Collection(dbName)

	filter := bson.M{
		"userInfo.email": value,
	}

	res := collection.FindOne(rep.cntx.cntx, filter)

	if res.Err() != nil {
		return nil, res.Err()
	}

	var result reg.CheckStatusResponseModel

	err = res.Decode(&result)

	if err != nil {
		return nil, err
	}

	return &result, err
}
