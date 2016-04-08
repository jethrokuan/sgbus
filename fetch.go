package main

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

var (
	AccountKey   string
	UniqueUserID string
	busURL       string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	AccountKey = os.Getenv("AccountKey")
	UniqueUserID = os.Getenv("UniqueUserID")
}

func makeBusStopURL(serviceNo string) string {
	return fmt.Sprintf("http://datamall2.mytransport.sg/ltaodataservice/BusArrival?BusStopID=%s&SST=True", serviceNo)
}

func getJson(url string, target interface{}) error {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Println(err)
		return err
	}

	req.Header.Set("AccountKey", AccountKey)
	req.Header.Set("UniqueUserID", UniqueUserID)
	req.Header.Set("accept", "application/json")

	res, err := client.Do(req)

	if err != nil {
		log.Println(err)
		return err
	}
	defer res.Body.Close()

	return json.NewDecoder(res.Body).Decode(target)
}
