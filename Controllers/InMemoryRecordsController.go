package controllers

import (
	"encoding/json"
	"errors"
	models "go-service/models"
	"go-service/utils"
	"net/http"
)

var (
	globalMap = make(map[string]string)
)

func ProcessInMemoryRecords(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if req.Method == "GET" {
		params := req.URL.Query()
		key := params.Get("key")
		val, err := FetchInMemoryRecord(key)
		if err != nil {
			json.NewEncoder(w).Encode(err)
			return
		}
		json.NewEncoder(w).Encode(models.InMemoryRecordsRequest{Key: key, Value: val})
		return

	} else if req.Method == "POST" {
		var recordRequest models.InMemoryRecordsRequest

		defer req.Body.Close()
		err := json.NewDecoder(req.Body).Decode(&recordRequest)
		if err != nil {
			errResponse, _ := json.Marshal(utils.GetErrorObject("INVALIDREQUEST"))
			w.Write(errResponse)
			return
		}
		SaveRecordInMemory(recordRequest)
		json.NewEncoder(w).Encode(recordRequest)
		return
	} else {
		json.NewEncoder(w).Encode(utils.GetErrorObject("METHODNOTSUPPORTED"))
		return
	}
}

func FetchInMemoryRecord(key string) (string, error) {
	if _, ok := globalMap[key]; ok {
		return globalMap[key], nil
	} else {
		return "", errors.New("key not found")
	}
}

func SaveRecordInMemory(record models.InMemoryRecordsRequest) {
	if _, ok := globalMap[record.Key]; ok {
		globalMap[record.Key] = record.Value
	} else {
		globalMap[record.Key] = record.Value
	}
}
