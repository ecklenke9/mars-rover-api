package router

import (
	"github.com/gin-gonic/gin"
	"mars-rover-api/pkg/controller"
	"mars-rover-api/pkg/model"
	"net/http"
	"time"
)

// New create the routes
func New(ctlr *controller.Controller) {
	router := gin.Default()
	router.GET("/roverImages", getRoverImages(ctlr))
	router.NoRoute(func(c *gin.Context) {
		c.AbortWithStatus(http.StatusNotFound)
	})
	router.Run(":8080")
}

// getRoverImages is a Router function that calls the Controller to get the
// last 10 days of rover images
func getRoverImages(ctlr *controller.Controller) gin.HandlerFunc {
	return func(c *gin.Context) {
		days, err := ctlr.GetRoverImages(GetDates())
		if err != nil {
			// Error bubbled up from either the controller.Client or controller.Mongo
			c.JSON(http.StatusBadRequest, err)
		}

		// Controller processed the request successfully
		c.JSON(http.StatusOK, Respond(days))
	}
}

// Respond is a Router helper function that takes []model.Day
func Respond(days []model.Day) map[string][]string {
	resp := map[string][]string{}
	for _, day := range days {
		if day.Images == nil {
			day.Images = []string{}
		}
		resp[day.Date] = day.Images
	}

	return resp
}

func GetDates() []model.Day {
	t := time.Now()

	dayArray := make([]model.Day, 0)
	// Pass in 10 days
	// so this loop stops at 10
	for i := 0; i < 10; i++ {
		// Format the date into YYYY-MM-DD
		date := t.Format("2006-01-02")
		// Add date to array of strings being returned
		dayArray = append(dayArray, model.Day{
			Date: date,
		})
		// Subtract 24 hours from the time object
		t = t.AddDate(0, 0, -1)
	}

	return dayArray
}
