package reps

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FeedbackMessageModel struct {
	UserID primitive.ObjectID
	Image  string
	Name   string
	Date   time.Time
	IsMe   bool
	Text   string `json:"text"`
}

type CreateFeedbackModel struct {
	UserId primitive.ObjectID
	Text   string   `json:"text"`
	Audio  string   `json:"audio"`
	Images []string `json:"images"`
}

type FeedbackDockModel struct {
	ID       primitive.ObjectID `bson:"_id"`
	UserID   primitive.ObjectID
	Audio    string
	Text     string
	Images   []string
	Comments *[]MessageModel
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
func InitFeedbackRep(cntx *DBContext) *FeedbackRepo {
	return &FeedbackRepo{cntx: cntx}
}

func (rep *FeedbackRepo) AddComment(msg *FeedbackMessageModel, id string) error {

	err := rep.cntx.client.Ping(rep.cntx.cntx, nil)

	collection := rep.cntx.db.Collection(feedbackDBName)

	objID, _ := primitive.ObjectIDFromHex(id)

	filer := bson.M{
		"_id": objID,
	}

	message := bson.M{
		"$push": bson.M{
			"comments": msg,
		},
	}

	prmid, _ := primitive.ObjectIDFromHex(id)

	t := collection.FindOne(rep.cntx.cntx, bson.M{"_id": prmid})

	if t.Err() != nil {
		return t.Err()
	}

	var tml FeedbackDockModel

	t.Decode(&tml)

	if tml.UserID.Hex() == msg.UserID.Hex() {
		msg.IsMe = true
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

	if err != nil {
		return err
	}

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

	var docks []FeedbackDockModel

	err = res.All(rep.cntx.cntx, &docks)

	if err != nil {
		return nil, err
	}

	return &docks, nil
}
