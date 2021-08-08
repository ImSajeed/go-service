package models

import (
	"strings"
	"time"
)

const dateFormat = "2006-01-02"

type MongoRecordsRequest struct {
	StartDate CustomTime `json:"startDate"`
	EndDate   CustomTime `json:"endDate"`
	MinCount  int        `json:"minCount"`
	MaxCount  int        `json:"maxCount"`
}

type MongoRecordsResponse struct {
	Code    int       `json:"code"`
	Msg     string    `json:"msg"`
	Records []Records `json:"records"`
}

type Records struct {
	Key        string    `bson:"key"`
	CreatedAt  time.Time `bson:"createdAt"`
	TotalCount int       `bson:"totalCount,omitempty"`
}

type CustomTime struct {
	time.Time
}

func (c *CustomTime) UnmarshalJSON(p []byte) error {
	t, err := time.Parse(dateFormat, strings.Replace(
		string(p),
		"\"",
		"",
		-1,
	))

	if err != nil {
		return err
	}

	c.Time = t

	return nil
}
