package constants

const API_SUCCESS_STATUS = "Success"

const API_FAILED_STATUS = "Failed"

const ApiFailStatus = "Fail"

var INVALID_TOKEN_RESPONSE = map[string]interface{}{
	"status":  ApiFailStatus,
	"message": "Invalid or No Token",
}
