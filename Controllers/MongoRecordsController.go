package Controllers

import (
	"context"
	"encoding/json"
	models "go-service/Models"
	utils "go-service/Utils"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

//Connection mongoDB with utils module
var collection = utils.ConnectDB()

func FetchMongoRecords(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var recordRequest models.MongoRecordsRequest
	var recordResponse models.MongoRecordsResponse

	defer req.Body.Close()
	err := json.NewDecoder(req.Body).Decode(&recordRequest)
	if err != nil {
		errorObject := utils.GetErrorObject("INVALIDREEQUEST")
		json.NewEncoder(w).Encode(models.MongoRecordsResponse{Code: errorObject.ErrorCode, Msg: errorObject.ErrorMsg, Records: []models.Records{}})
		return
	}

	// creating mongo equivalent query
	project := bson.D{primitive.E{Key: "$project", Value: bson.D{primitive.E{Key: "key", Value: "$key"}, primitive.E{Key: "createdAt", Value: "$createdAt"},
		primitive.E{Key: "totalCount", Value: bson.D{primitive.E{Key: "$sum", Value: "$counts"}}}},
	}}

	match := bson.D{primitive.E{Key: "$match", Value: bson.D{primitive.E{Key: "createdAt", Value: bson.D{primitive.E{Key: "$gte", Value: primitive.NewDateTimeFromTime(recordRequest.StartDate.Time)}, primitive.E{Key: "$lte", Value: primitive.NewDateTimeFromTime(recordRequest.EndDate.Time)}}}}}}

	match2 := bson.D{primitive.E{Key: "$match", Value: bson.D{primitive.E{Key: "totalCount", Value: bson.D{primitive.E{Key: "$gte", Value: recordRequest.MinCount}, primitive.E{Key: "$lte", Value: recordRequest.MaxCount}}}}}}

	cur, err := collection.Aggregate(context.TODO(), mongo.Pipeline{project, match, match2})

	if err != nil {
		errorObject := utils.GetErrorObject("INTERNALERROR")
		json.NewEncoder(w).Encode(models.MongoRecordsResponse{Code: errorObject.ErrorCode, Msg: errorObject.ErrorMsg, Records: []models.Records{}})
		return
	}

	var records []models.Records
	if err = cur.All(context.TODO(), &records); err != nil {
		errorObject := utils.GetErrorObject("INTERNALERROR")
		json.NewEncoder(w).Encode(models.MongoRecordsResponse{Code: errorObject.ErrorCode, Msg: errorObject.ErrorMsg, Records: []models.Records{}})
		return
	}

	recordResponse.Code = 0
	recordResponse.Msg = "success"
	recordResponse.Records = records

	json.NewEncoder(w).Encode(recordResponse)

}
