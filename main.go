package main

import "encoding/json"
import "fmt"
import "net/http"
import "os"
import "github.com/joho/godotenv"

type Property struct {
	Value float32 `json:"value"`
}

type Properties struct {
	Summary     string   `json:"textDescription"`
	Humidity    Property `json:"relativeHumidity"`
	Temperature Property `json:"temperature"`
}

type Weather struct {
	Properties Properties `json:"properties"`
}

func main() {
	godotenv.Load()

	var params []string
	params = append(params, "apiKey="+os.Getenv("WX_API_KEY"))
	params = append(params, "geocode="+os.Getenv("WX_GEO_LAT")+"%2C"+os.Getenv("WX_GEO_LONG"))
	params = append(params, "units=e&language=en-US&format=json")

	url := "https://api.weather.gov/stations/" + os.Args[1] + "/observations/latest"
	res, _ := http.Get(url)

	var weather *Weather
	json.NewDecoder(res.Body).Decode(&weather)

	fmt.Printf("%.1f° / %.1f%% - %v", toF(weather.Properties.Temperature.Value), weather.Properties.Humidity.Value, weather.Properties.Summary)
}

func toF(cVal float32) float32 {
	return (cVal * 9 / 5) + 32
}
