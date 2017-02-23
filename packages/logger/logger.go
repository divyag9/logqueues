package logger

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	"github.com/divyag9/goqueues/packages/queue"
)

// LogQueueDataToFile appaneds/writes the queue data to the target file
func LogQueueDataToFile(data []byte, targetFile string) error {
	err := createIfFileNotExists(targetFile)
	if err != nil {
		return err
	}
	queueData, err := parseData(data)
	if err != nil {
		return err
	}
	err = writeFile(queueData, targetFile)
	if err != nil {
		return err
	}

	return nil
}

func createIfFileNotExists(targetFile string) error {
	if _, err := os.Stat(targetFile); os.IsNotExist(err) {
		file, err := os.Create(targetFile)
		if err != nil {
			return fmt.Errorf("Unable to create target file: %s. Error: %s", targetFile, err)
		}
		file.Close()
	}
	return nil
}

func parseData(data []byte) ([]queue.Details, error) {
	var queueData []queue.Details
	err := json.Unmarshal(data, &queueData)
	if err != nil {
		return nil, fmt.Errorf("Unable to unmarshal queue data: %s", err)
	}
	return queueData, nil
}

func writeFile(queueData []queue.Details, targetFile string) error {
	fo, err := os.OpenFile(targetFile, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return fmt.Errorf("Error opening target file: %s", err)
	}
	defer fo.Close()

	bufferedWriter := bufio.NewWriter(fo)
	for _, queue := range queueData {
		logValue := fmt.Sprintf("Name=%v, Type=%v, Depth=%v, Rate=%v, LastProcessed=%v, LastReported=%v", queue.Name, queue.Type, queue.Depth, queue.Rate, queue.LastProcessed, queue.LastReported)
		_, err := fmt.Fprintln(bufferedWriter, logValue)
		if err != nil {
			return fmt.Errorf("Error writing to target file: %s", err)
		}
	}
	bufferedWriter.Flush()
	return nil
}
