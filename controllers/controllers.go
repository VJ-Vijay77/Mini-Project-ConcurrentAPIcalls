package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/go-resty/resty/v2"
)

type APIresponse struct {
	endpoint string
	data     string
	err      error
}

func HandleApi1(w http.ResponseWriter, r *http.Request) {
	Data := "Data from API 1"
	sendJSON(w, Data)
}

func HandleApi2(w http.ResponseWriter, r *http.Request) {
	Data := "Data from API 2"
	sendJSON(w, Data)

}

func HandleApi3(w http.ResponseWriter, r *http.Request) {
	Data := "Data from API 3"
	sendJSON(w, Data)

}

func sendJSON(w http.ResponseWriter, data string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Error marshalling json", http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
}

func ConcurrentAPI(w http.ResponseWriter, r *http.Request) {
	apiEndPoints := []string{"http://localhost:8080/api-data-1", "invalid://api-endpoint", "http://localhost:8080/api-data-2", "http://localhost:8080/api-data-3"}

	var wg sync.WaitGroup
	resultChannel := make(chan APIresponse, len(apiEndPoints))

	for _, v := range apiEndPoints {
		wg.Add(1)
		go fetchAPIs(v, &wg, resultChannel)
	}

	go func() {
		wg.Wait()
		close(resultChannel)
	}()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	for v := range resultChannel {
		if v.err != nil {
			processedError := processError(v.endpoint, v.err)
			processedError = append(processedError, '\n')
			w.Write(processedError)
			continue
		}
		data, _ := json.Marshal(fmt.Sprintf("Response from API :%s Data:%s", v.endpoint, v.data))
		data = append(data, '\n')
		w.Write(data)
	}

}

func processError(endpoint string, err error) []byte {
	data, _ := json.Marshal(fmt.Sprintf("Error getting response from endpoint :%s with error %v", endpoint, err))
	return data

}

func fetchAPIs(endpoint string, wg *sync.WaitGroup, resultChannel chan<- APIresponse) {
	defer wg.Done()

	client := resty.New()

	resp, err := client.R().Get(endpoint)
	if err != nil {
		resultChannel <- APIresponse{endpoint: endpoint, err: err}
		return
	}

	var dataFromRequest string

	err = json.Unmarshal(resp.Body(), &dataFromRequest)
	if err != nil {
		resultChannel <- APIresponse{endpoint: endpoint, err: err}
		return
	}

	resultChannel <- APIresponse{endpoint: endpoint, data: dataFromRequest}

}
