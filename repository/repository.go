package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/anujsharma13/model"
	"github.com/anujsharma13/redis"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://Anuj:Anuj123@cluster0.neuwedn.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"
const dbName = "sahayak"
const collectionName = "workers"

var collection *mongo.Collection

func init() {
	clientoption := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.TODO(), clientoption)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Success Mongo connection")
	collection = client.Database(dbName).Collection(collectionName)

	fmt.Println("Mongo connection success")
}

func InsertOneWorker(worker model.Workers) {
	inserted, err := collection.InsertOne(context.Background(), worker)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted one worker in Db", inserted.InsertedID)
}

func UpdateOneWorker(workerId string) {
	// id, err := primitive.ObjectIDFromHex(workerId)
	// if err != nil {
	// 	panic(err)
	// }
	// // THERE are three bson.A, bson.M (ordered), bson.D each are different
	// filter := bson.M{"_id": id}
	// update := bson.M{"$set": bson.M{"watched": true}}
	// result, err := collection.UpdateOne(context.Background(), filter, update)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Modified count", result.ModifiedCount)
}

func DeleteOneWorker(workerId string) {
	id, _ := primitive.ObjectIDFromHex(workerId)
	filter := bson.M{"_id": id}
	result, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("total deleted", result)
}

func GetAllWorkers() []primitive.M {
	filter := bson.D{{}}
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	var workers []primitive.M
	for cursor.Next(context.Background()) {
		var worker bson.M
		err := cursor.Decode(&worker)
		if err != nil {
			log.Fatal(err)
		}
		workers = append(workers, worker)
	}
	defer cursor.Close(context.Background())
	return workers
}

func GetOneWorker(workerId string) (model.Workers, error) {
	worker, err := redis.Get(workerId)
	if err != nil {
		return model.Workers{}, err
	}
	filter := bson.D{{"_id", workerId}}
	var worker2 model.Workers
	err1 := collection.FindOne(context.Background(), filter).Decode(&worker2)
	if err1 != nil {
		return model.Workers{}, err
	}
	return worker, nil
}
