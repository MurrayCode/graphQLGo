package database

import (
	"context"
	"log"
	"time"

	"github.com/MurrayCode/graphQLGo/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	client *mongo.Client
}

func Connect() *DB {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return &DB{
		client: client,
	}
}

//Add New Watch
func (db *DB) Save(input *model.NewWatch) *model.Watch {
	collection := db.client.Database("watchstore").Collection("watches")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := collection.InsertOne(ctx, input)
	if err != nil {
		log.Fatal(err)
	}
	return &model.Watch{
		ID:      res.InsertedID.(primitive.ObjectID).Hex(),
		Name:    input.Name,
		Brand:   input.Brand,
		Price:   input.Price,
		Stock:   input.Stock,
		InStock: input.InStock,
	}
}

//Delete Watch by ID
func (db *DB) Delete(ID string) *model.Watch {
	ObjectID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		log.Fatal(err)
	}
	collection := db.client.Database("watchstore").Collection("watches")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res := collection.FindOneAndDelete(ctx, bson.M{"_id": ObjectID})
	watch := model.Watch{}
	res.Decode(&watch)
	return &watch
}

//Find Watch by ID
func (db *DB) FindByID(ID string) *model.Watch {
	ObjectID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		log.Fatal(err)
	}
	collection := db.client.Database("watchstore").Collection("watches")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res := collection.FindOne(ctx, bson.M{"_id": ObjectID})
	watch := model.Watch{}
	res.Decode(&watch)
	return &watch
}

//Update Watch by ID
func (db *DB) Update(ID string, input *model.UpdateWatch) *model.Watch {
	ObjectID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		log.Fatal(err)
	}
	collection := db.client.Database("watchstore").Collection("watches")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	update := bson.D{{"$set",
		bson.D{
			{"name", input.Name},
			{"brand", input.Brand},
			{"price", input.Price},
			{"stock", input.Stock},
			{"inStock", input.InStock},
		}}}
	res := collection.FindOneAndUpdate(ctx, bson.M{"_id": ObjectID}, update)
	watch := model.Watch{}
	res.Decode(&watch)
	return &watch
}

//Get All Watches
func (db *DB) All() []*model.Watch {
	collection := db.client.Database("watchstore").Collection("watches")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	var watches []*model.Watch
	for cur.Next(ctx) {
		var watch *model.Watch
		err := cur.Decode(&watch)
		if err != nil {
			log.Fatal(err)
		}
		watches = append(watches, watch)
	}
	return watches
}
