package model

type NasaResponse struct {
	Photos []Photos `json:"photos"`
}
type Camera struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	RoverID  int    `json:"roverId"`
	FullName string `json:"fullName"`
}
type Rover struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	LandingDate string `json:"landingDate"`
	LaunchDate  string `json:"launchDate"`
	Status      string `json:"status"`
}
type Photos struct {
	ID        int    `json:"id"`
	Sol       int    `json:"sol"`
	Camera    Camera `json:"camera"`
	ImgSrc    string `json:"imgSrc"`
	EarthDate string `json:"earthDate"`
	Rover     Rover  `json:"rover"`
}

type Day struct {
	Date   string   `bson:"date"`
	Images []string `bson:"images"`
}

type ApiResponse struct {
	TenDayArray map[string][]string `bson:"tenDayArray"`
}
