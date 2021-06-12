package httpclient

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"io/ioutil"
	"log"
	"mars-rover-api/pkg/model"
	"net/http"
	"os"
)

// Constant nasaUrl to hold Nasa url and query params
const nasaUrl = "https://api.nasa.gov/mars-photos/api/v1/rovers/curiosity/photos?earth_date=%s&api_key=%s"

type Client struct {
	client *http.Client
	ApiKey string
}

// logError will log any errors that are returned from functions
func logError(f string, err error) {
	if err != nil {
		log.Printf("function: %v; error: %v", f, err)
	}
}

// Connect loads in the .env file and assigns environment variables
func Connect() (*Client, error) {
	var err error
	defer logError("Connect()", err)

	err = godotenv.Load(".env")
	if err != nil {
		return nil, err
	}

	apiKey := os.Getenv("NASA_API_KEY")

	cli := http.Client{
		Transport: http.DefaultTransport,
		Timeout:   5000,
	}

	return &Client{client: &cli, ApiKey: apiKey}, nil
}

// Call CallNasaApi is a method of the Client object that will
// call the Nasa API to retrieve the last 10 days of rover images
func (c *Client) CallNasaApi(earthDate string) ([]string, error) {
	var err error
	defer logError("CallNasaApi()", err)

	// Create a new request with the Nasa url that contains the query params "earth_date", and "api_key"
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf(nasaUrl, earthDate, c.ApiKey), nil)
	if err != nil {
		return nil, err
	}

	// Add Header to request
	req.Header.Add("Accept", "application/json")

	// Send an http.Request that will return an http.Response
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	// Read the response body into a []byte
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	// Unmarshall body into new nasaResponse
	var nasaResponse model.NasaClientResponse
	err = json.Unmarshal(body, &nasaResponse)
	if err != nil {
		return nil, err
	}

	// Create []respImages that will return the 3 images
	var respImages []string
	respImages = processNasaResponse(nasaResponse)
	return respImages, nil
}

// processNasaResponse will iterate over nasaResponse to grab 3 images
func processNasaResponse(nasaResponse model.NasaClientResponse) []string {
	imageArray := make([]string, 0)

	// If Photos do not exist
	if len(nasaResponse.Photos) == 0 {
		// Return the empty imageArray
		return imageArray
	}

	// For every Photo in the response from nasa
	for _, p := range nasaResponse.Photos {
		// Append the imgSrc to the imageArry
		imageArray = append(imageArray, p.ImgSrc)
		// If imageArray is 3 photos full
		if len(imageArray) == 3 {
			// Return imageArray
			return imageArray
		}
	}

	return imageArray
}
