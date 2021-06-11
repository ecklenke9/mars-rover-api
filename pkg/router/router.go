package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mars-rover-api/pkg/controller"
	"net/http"
	"time"
)

//New function
func New(ctlr controller.Controller) {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.GET("/roverImages", getRoverImages(ctlr))
	}
	router.NoRoute(func(c *gin.Context) {
		c.AbortWithStatus(http.StatusNotFound)
	})
	router.Run(":8080")
}

func getRoverImages(ctlr controller.Controller) gin.HandlerFunc {
	return func(c *gin.Context) {

		roverImageArray, err := ctlr.GetRoverImages(getDates())
		if err != nil {
			// error bubbled up from either the controller.Client or controller.Mongo
			c.JSON(http.StatusBadRequest, err)
		}

		// controller processed the request successfully
		c.JSON(http.StatusOK, roverImageArray)
	}
}

func getDates() []string {
	t := time.Now()
	dateArray := make([]string, 0)

	// we need the past ten days,
	// so this loop stops at ten
	for i := 0; i < 10; i++ {
		// format the date into YYYY-MM-DD
		date := t.Format("2006-01-02")
		// add date to array of strings being returned
		dateArray = append(dateArray, date)
		// subtract 24 hours from the time object
		t = t.AddDate(0, 0, -1)
		fmt.Println("Added date to dateArray: ", date)
	}

	return dateArray
}
