package reps

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MessageModel struct {
	Author string `json:"author"`
	Text   string `json:"text"`
}

type CreateFeedbackModel struct {
	Filepath *string `json:"аilepath"`
	Text     string  `json:"text"`
}

type FeedbackDockModel struct {
	ID string `json:"id"`

	Filepath *string `json:"аilepath"`
	Text     string  `json:"text"`

	Messages *[]MessageModel `json:"text"`
}

type NotFound struct {
}

func (er *NotFound) Error() string {
	return "Запись с таким ID не найдена"
}

const feedbackDBName = "feedback"

// FeedbackRepo репозиторий регистрации
type FeedbackRepo struct {
	cntx *DBContext
}

// InitFeedbackRep иницаллизирует репозиторий
func InitFeedbackRep(cntx *DBContext) *RegRep {
	return &RegRep{cntx: cntx}
}

func (rep *FeedbackRepo) AddComment(msg *MessageModel, id string) error {

	err := rep.cntx.client.Ping(rep.cntx.cntx, nil)

	collection := rep.cntx.db.Collection(feedbackDBName)

	objID, _ := primitive.ObjectIDFromHex(id)

	filer := bson.M{
		"_id": objID,
	}

	message := bson.M{
		"$push": bson.M{
			"comments": bson.M{
				"author": msg.Author,
				"text":   msg.Text,
			},
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

func (rep *FeedbackRepo) CreatePost(post *CreateFeedbackModel) error {
	err := rep.cntx.client.Ping(rep.cntx.cntx, nil)

	collection := rep.cntx.db.Collection(feedbackDBName)

	_, err = collection.InsertOne(rep.cntx.cntx, post)

	if err != nil {
		return err
	}

	return nil
}

func (rep *FeedbackRepo) GetAllPosts() (*[]FeedbackDockModel, error) {
	err := rep.cntx.client.Ping(rep.cntx.cntx, nil)

	collection := rep.cntx.db.Collection(feedbackDBName)

	res, err := collection.Find(rep.cntx.cntx, bson.M{})

	if err != nil {
		return nil, err
	}

	var docks *[]FeedbackDockModel

	err = res.All(rep.cntx.cntx, docks)

	if err != nil {
		return nil, err
	}

	return docks, nil
}
