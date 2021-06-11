package httpclient

import (
	"fmt"
	"github.com/joho/godotenv"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Client struct {
	client *http.Client
	ApiKey string
}

//func main(){
//	cli, err := New()
//	if err != nil{
//		log.Fatalf("error loading .env file")
//	}
//	cli.callNasaApi()
//}

func New() (*Client, error) {
	err := godotenv.Load(".env")
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

func (c *Client) CallNasaApi(date string) ([]string, error) {
	req, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("https://api.nasa.gov/mars-photos/api/v1/rovers/curiosity/photos?api_key=%s", c.ApiKey),
		nil,
	)
	if err != nil {
		log.Fatalf("error creating HTTP request: %v", err)
	}

	req.Header.Add("Accept", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("error sending HTTP request: %v", err)
	}
	responseBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("error reading HTTP response body: %v", err)
	}

	log.Println("We got the response:", string(responseBytes))

	// TODO: Add return values
	return nil, nil
}
