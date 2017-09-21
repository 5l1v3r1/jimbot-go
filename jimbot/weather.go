package jimbot

import (
	"fmt"
	"github.com/buger/jsonparser"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const (
	clouds = "⛅"
	sunny  = "☀️"
	rain   = "🌧"
	snow   = "🌨"
)

// NowWeather : Current weather info from his/her city
func NowWeather(city string) string {
	apiURL := fmt.Sprintf("https://api.seniverse.com/v3/weather/now.json?key=jkfkmayhcqnvsrjr&location=%s&language=en&unit=c", city)

	resp, err := http.Get(apiURL)
	if err != nil {
		log.Print("[===] WEATHER API cant be reached")
	}
	defer resp.Body.Close()

	readBody, _ := ioutil.ReadAll(resp.Body)

	status, _, _, _ := jsonparser.Get(readBody, "results", "[0]", "now", "text")
	temp, _, _, _ := jsonparser.Get(readBody, "results", "[0]", "now", "temperature")
	var retVal string
	retVal = "`Now in " + strings.ToUpper(city) + ": " + string(status) + "\nTemp: " + string(temp) + " °C`"
	return retVal
}
