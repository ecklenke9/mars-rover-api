package main

import (
	"log"
	"mars-rover-api/pkg/controller"
	"mars-rover-api/pkg/router"
)

func main() {
	// Create new controller to handle Nasa client and Mongo client
	ctlr, err := controller.New()
	if err != nil {
		log.Fatal(err)
	}

	// Uncomment code "router.New(ctlr)" to pass controller to router to run this
	// application as a web api hosted on a server
	// router.New(ctlr)

	// Call GetRoverImages to get the Rover images ([]model.Day)
	images, err := ctlr.GetRoverImages(router.GetDates())
	if err != nil {
		// Error bubbled up from either the controller.Client or controller.Mongo
		log.Println(err)
	}

	// Print the Rover images to console
	log.Println(images)
}
