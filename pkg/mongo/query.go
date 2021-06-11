package mongo

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"mars-rover-api/pkg/model"
)

func (db Mongo) CheckCache(dates []string) ([]model.Day, error) {
	// create an array of day objects
	days := make([]model.Day, 0)
	for _, d := range dates {
		days = append(days, model.Day{
			Date: d,
		})
	}
	// mongo.query where you're searching for all entries in
	//client.yourCollectionName.find({_id:{$in:[yourValue1,yourValue2,yourValue3…
	query := bson.D{{"date", bson.D{{"$in", dates}}}}
	results, err := db.client.Database("nasaDB").Collection(RoverCollection).Find(context.Background(), query)
	if err != nil || results == nil {
		if err == nil {
			return nil, errors.New("mongo returned nil results")
		}
		return nil, err
	}

	err = results.All(context.Background(), &days)
	if err != nil {
		return nil, err
	}
	return days, nil
}

func (db Mongo) UpsertRoverImages(dates []model.Day) ([]model.Rover, error) {

	//query := bson.D{{"date", bson.D{{"$in", dates}}}}

	// TODO: Work on Upsert to Mongo Functionality
	//results, err := db.client.Database("nasaDB").Collection(RoverCollection).UpdateMany(context.Background(), query)
	//if err != nil || results == nil{
	//	if err == nil{
	//		return nil, errors.New("mongo returned nil results")
	//	}
	//	return nil, err
	//}
	//
	//err = results.All(context.Background(), &days)
	//if err != nil{
	//	return nil, err
	//}

	// TODO: Add return values
	return nil, nil
}
