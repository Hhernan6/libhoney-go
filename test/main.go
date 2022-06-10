package main

import (
	"log"
	"net/http"
	"time"

	libhoney "github.com/honeycombio/libhoney-go"
	"github.com/honeycombio/libhoney-go/marker"
)

func main() {
	err := libhoney.Init(libhoney.Config{
		WriteKey: "<API_KEY>",
		Dataset:  "MyCodeExerciseDataset",
		Marker: marker.Sender{
			HTTPClient: &http.Client{Timeout: 30 * time.Second},
		},
	})
	if err != nil {
		log.Println("error iniitalizing libhoney library: %w", err)
		panic(err)
	}

	defer libhoney.Close()
	markerClient := libhoney.NewMarker()
	data := marker.CreateMarkerData{
		Message: "deploy #344",
		Type:    "testing three",
	}

	response, err := markerClient.SetMarker(data)
	if err != nil {
		log.Println("error creating marker: %w", err)
	}

	log.Println(response)
}
