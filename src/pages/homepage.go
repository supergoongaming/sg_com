package pages

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

type homepageTemplateData struct {
	Year string
}

func getCurrentYear() string {
	currentTime := time.Now()
	return strconv.Itoa(currentTime.Year())
}

func HomepageHandler(w http.ResponseWriter, r *http.Request) {
	data := homepageTemplateData{
		Year: getCurrentYear(),
	}
	err := GetTemplates().ExecuteTemplate(w, "index.html", data)
	if err != nil {
		log.Printf("Error when loading template:\n %s", err.Error())
		fmt.Fprintf(w, "Failed to load template content properly.")
	}
}
