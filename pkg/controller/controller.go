package controller

import (
	"mars-rover-api/pkg/httpclient"
	"mars-rover-api/pkg/model"
	"mars-rover-api/pkg/mongo"
)

// Controller - struct - contains a pointer to a Mongo and Client object
// This object will be used in the router function so that each of the routes
// has access to everything it needs in the application in order to perform
type Controller struct {
	Client *httpclient.Client
	Mongo  *mongo.Mongo
}

// New() - func - calls httpclient.New() and mongo.New()
// in order to build out the client communications required for processing
func New() (*Controller, error) {
	// build httpclient.Client
	client, err := httpclient.New()
	if err != nil {
		return nil, err
	}

	// build mongo.Mongo
	db, err := mongo.Connect()
	if err != nil {
		return nil, err
	}

	// assign the mongo and client connections to a controller object and return
	return &Controller{Client: client, Mongo: db}, nil
}

// GetRoverImages Endpoint
func (ctlr *Controller) GetRoverImages(dates []string) ([]model.Rover, error) {
	// attempt to get the rover image array from Mongo first
	daysArray, err := ctlr.Mongo.CheckCache(dates)
	if err != nil {
		return nil, err
	}

	upsertArray := make([]model.Day, 0)
	// loop over mongo results
	for _, d := range daysArray {
		// if the images for the day are empty
		var imageArray []string
		if d.Images == nil {
			// attempt to reach out to NASA for images
			imageArray, err = ctlr.Client.CallNasaApi(d.Date)
			if err != nil {
				// could not find images
				// setting d.Images to an array of strings
				// instead of nil as per the Exercise Text
				// then continue to process the next mongo result
				d.Images = []string{}
				continue
			}

			// attach found images to the day(d)
			// store this day in mongo for
			// the next call to this endpoint
			d.Images = imageArray
			upsertArray = append(upsertArray, d)
		}
	}

	upsertResult, err := ctlr.Mongo.UpsertRoverImages(upsertArray)
	if err != nil {
		return nil, err
	}

	return upsertResult, nil
}
