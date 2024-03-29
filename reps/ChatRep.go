package reps

import (
	"suncity/commod"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ChatRep struct {
	cntx *DBContext
}

func InitChatRep(cntx *DBContext) *ChatRep {
	return &ChatRep{cntx: cntx}
}

func (rep *ChatRep) SaveMessage(msg *commod.MessageModel, user *commod.ServiceUser) error {
	err := rep.cntx.client.Ping(rep.cntx.cntx, nil)

	collection := rep.cntx.db.Collection(dbName)

	filer := bson.M{
		"_id": user.ID,
	}

	message := bson.M{
		"$push": bson.M{
			"messages": msg,
		},
	}

	res, err := collection.UpdateOne(rep.cntx.cntx, filer, message)

	if err != nil {
		return err
	}

	if res.ModifiedCount == 0 {
		return &NotFound{}
	}

	return nil
}

func (rep *ChatRep) SaveMessageById(msg *commod.MessageModel, userId string) error {
	err := rep.cntx.client.Ping(rep.cntx.cntx, nil)

	collection := rep.cntx.db.Collection(dbName)

	objid, _ := primitive.ObjectIDFromHex(userId)

	filer := bson.M{
		"_id": objid,
	}

	message := bson.M{
		"$push": bson.M{
			"messages": msg,
		},
	}

	res, err := collection.UpdateOne(rep.cntx.cntx, filer, message)

	if err != nil {
		return err
	}

	if res.ModifiedCount == 0 {
		return &NotFound{}
	}

	return nil
}
