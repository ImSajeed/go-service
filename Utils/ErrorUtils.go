package utils

type ErrorObject struct {
	ErrorCode int
	ErrorMsg  string
}

var (
	errorsMap = map[string]ErrorObject{
		"INVALIDREEQUEST": {ErrorCode: 1, ErrorMsg: "Invalid Request"},
		"INTERNALERROR":  {ErrorCode: 2, ErrorMsg: "Internal Server Error"},
		"METHODNOTSUPPORTED":  {ErrorCode: 3, ErrorMsg: "Method Not Supported"},
	}
)

func GetErrorObject(name string) ErrorObject {
	if _, ok := errorsMap[name]; ok {
		return errorsMap[name]
	} else {
		return ErrorObject{ErrorCode: 4, ErrorMsg: ""}
	}
}
