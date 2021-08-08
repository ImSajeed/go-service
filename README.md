# go-service

This application consists of 3 endpoints built on Go basic http server

1) POST /fetchMongoRecords - fetch the records from Mongo DB for given request details {startDate,endDate,minCount and maxCount}
2) POST /in-memory - saves the given record(key-value pair) in-memory using go map (updates if same key exists)
3) GET /in-memorry?key=value - fetches the records stored in-memory go map based on given key

cURLS:

POST /fetchMongoRecords

curl --location --request POST 'http://localhost:8090/fetchMongoRecords' \
--header 'Content-Type: application/json' \
--data-raw '{
    "startDate":"2016-01-26",
    "endDate":"2018-02-02",
    "minCount":2700,
    "maxCount":3000
}'

POST /in-memory

curl --location --request POST 'http://localhost:8090/in-memory' \
--header 'Content-Type: application/json' \
--data-raw '{
    "key":"active-tabs",
    "value":"get-go"
}'

GET /in-memorry?key=value

curl --location --request GET 'https://go-poc.herokuapp.com/in-memory?key=active-tabs'
