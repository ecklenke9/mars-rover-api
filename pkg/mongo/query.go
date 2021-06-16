package mongo

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"mars-rover-api/pkg/model"
)

func (db Mongo) CheckCache(initialDays []model.Day) ([]model.Day, error) {
	// For each date(d) in range dates
	datesForQuery := make([]string, 0)

	for _, d := range initialDays {
		// Create a new model.Day object
		// using the date(d) to populate the Date field
		// on the new model.Day object
		datesForQuery = append(datesForQuery, d.Date)
	}

	query := bson.D{{"date", bson.D{{"$in", datesForQuery}}}}
	results, err := db.client.Database("myFirstDatabase").Collection(Collection).Find(context.Background(), query)
	if err != nil || results == nil {
		if err == nil {
			return nil, errors.New("mongo returned nil results")
		}
		return nil, err
	}

	// Create an array of Day objects
	mongoResults := make([]model.Day, 0)
	err = results.All(context.Background(), &mongoResults)
	if err != nil {
		return nil, err
	}

	finalResults := captureFoundImages(initialDays, mongoResults)

	return finalResults, nil
}

func captureFoundImages(initialDays, mongoResults []model.Day) []model.Day {
	// For each day in initialDays
	for i, v := range initialDays {
		// Loop over the resultsDayArray to find a match
		// for each day in the resultsDayArray
		for _, d := range mongoResults {
			// If original.Date matches result.Date
			if v.Date == d.Date {
				// Move the images from the result(d) to the og(v)
				initialDays[i].Images = d.Images
			}
		}
	}

	// Return result
	return initialDays
}

func (db Mongo) UpsertRoverImages(upsertToMongo []model.Day) error {
	ops := buildWriteOperations(upsertToMongo)

	falsePtr := false
	_, err := db.client.Database("myFirstDatabase").Collection(Collection).BulkWrite(context.TODO(), ops, &options.BulkWriteOptions{Ordered: &falsePtr})
	if err != nil {
		return errors.New("unable to upsert images")
	}

	return nil
}

func buildWriteOperations(upsertToMongo []model.Day) []mongo.WriteModel {
	writes := make([]mongo.WriteModel, 0)
	for _, d := range upsertToMongo {
		op := mongo.NewUpdateOneModel()
		op.SetFilter(bson.M{"date": d.Date})
		op.SetUpdate(bson.M{"$set": bson.M{"images": d.Images}})
		op.SetUpsert(true)
		writes = append(writes, op)
	}

	return writes
}
