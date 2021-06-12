package controller

import (
	"log"
	"mars-rover-api/pkg/httpclient"
	"mars-rover-api/pkg/model"
	"mars-rover-api/pkg/mongo"
)

// The Controller struct contains a pointer to a Mongo and Client object
// This object will be used in the router function so that each of the routes
// has access to everything it needs within the application in order to perform
type Controller struct {
	Client *httpclient.Client
	Mongo  *mongo.Mongo
}

// New calls httpclient.Connect() and mongo.Connect()
// in order to build out the client communications required for processing
func New() (*Controller, error) {
	// Build *httpclient.Client
	client, err := httpclient.Connect()
	if err != nil {
		return nil, err
	}

	// Build *mongo.Mongo
	db, err := mongo.Connect()
	if err != nil {
		return nil, err
	}

	// Assign the Mongo and Client connections to a controller object and return
	return &Controller{Client: client, Mongo: db}, nil
}

// GetRoverImages attempts to get the Rover image array from Mongo first
func (ctlr *Controller) GetRoverImages(originalDaysArray []model.Day) ([]model.Day, error) {
	var err error
	originalDaysArray, err = ctlr.Mongo.CheckCache(originalDaysArray)
	if err != nil {
		return nil, err
	}

	// UpsertArray holds each day that hasn't been stored in Mongo
	upsertArray := make([]model.Day, 0)
	// Loop over mongo results
	for _, d := range originalDaysArray {
		// If the images for the day are empty
		var imageArray []string
		if d.Images == nil {
			// Attempt to reach out to NASA for images
			imageArray, err = ctlr.Client.CallNasaApi(d.Date)
			if err != nil {
				log.Printf("Error calling CallNasaApi: %v", err)
				// Could not find images
				// setting d.Images to an array of strings
				// instead of nil as per the Exercise Text
				// then continue to process the next mongo result
				d.Images = []string{}
				continue
			}
			// Attach found images to the day(d)
			// Store this day in mongo for
			// the next call to this endpoint
			d.Images = imageArray
			upsertArray = append(upsertArray, d)
		}
	}
	if len(upsertArray) != 0 {
		err = ctlr.Mongo.UpsertRoverImages(upsertArray)
		if err != nil {
			return nil, err
		}
	}

	return originalDaysArray, nil
}
