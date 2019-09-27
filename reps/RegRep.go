package reps

import (
	reg "suncity/reg/models"
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
