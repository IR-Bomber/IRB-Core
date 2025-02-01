package main

import (
	"fmt"
	"sync"

	"github.com/IR-Bomber/IRB-Go/bomber"
	"github.com/IR-Bomber/IRB-Go/utils"
)



func sendAllMessages(phoneNumber string) {
	fmt.Println("Sending SMS")

	// err, result := fetchAPI()
	err, result := utils.FetchAPI()



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

func main() {}

