package reps

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb://mongoadmin:secret@demo6.alpha.vkhackathon.com:27017/sun_city"
const dataBaseString = "sun_city"

const testConnectionString = "mongodb://mongoadmin:secret@demo6.alpha.vkhackathon.com:27017/sun_city"
const testDataBaseString = "test_sun_city"

// DBContext контекст для работы с БД.
type DBContext struct {
	client *mongo.Client
	db     *mongo.Database
	cntx   context.Context
}

// NewDB создает новый контекст и подключается к базе
func NewDB() (*DBContext, error) {
	return createDBContext(connectionString, dataBaseString)
}

// NewTestDB создает контекст для тестов и подключается к тестовой базе.
func NewTestDB() (*DBContext, error) {
	return createDBContext(testConnectionString, testDataBaseString)
}

// Close закрывает подключение к базе данных.
func (repo *DBContext) Close() error {
	return repo.client.Disconnect(repo.cntx)
}

func createDBContext(cnstr string, dbstr string) (*DBContext, error) {
	connectionClient := options.Client().ApplyURI(cnstr)
	client, err := mongo.NewClient(connectionClient)
	cntx := context.TODO()

	if err != nil {
		return nil, err
	}

	err = client.Connect(cntx)

	if err != nil {
		return nil, err
	}

	db := client.Database(dbstr)

	return &DBContext{client: client, db: db, cntx: cntx}, nil
}
