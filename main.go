package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/divyag9/logqueues/packages/logger"
)

func main() {
	// Parsing the command line arguments
	targetFile := flag.String("targetfile", "", "File to which the queue data needs to be logged")
	flag.Parse()

	if *targetFile == "" {
		log.Fatal("Please pass the required flags: -targetfile")
	}

	//Call goqueues and log every one minute
	ticker := time.NewTicker(60 * time.Second)
	for range ticker.C {
		getQueuesAndLog(*targetFile)
	}
}

func getQueuesAndLog(targetFile string) {
	queueData, err := getQueues()
	if err != nil {
		fmt.Println("Error calling goqueues: ", err)
	}
	err = logger.LogQueueDataToFile(queueData, targetFile)
	if err != nil {
		fmt.Println("Error logging queue data: ", err)
	}
}

func getQueues() ([]byte, error) {
	url := "http://localhost:8080/queues"
	var netClient = &http.Client{
		Timeout: time.Second * 10,
	}
	response, err := netClient.Get(url)
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("Error while reading the response: %s", err.Error())
	}
	return data, nil
}
