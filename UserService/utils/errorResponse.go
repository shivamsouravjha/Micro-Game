package utils

import "github.com/shivamsouravjha/Micro-Game/UserService/constants"

func SendErrorResponse(err error) interface{} {
	return map[string]interface{}{
		"status":  constants.API_FAILED_STATUS,
		"message": err.Error(),
	}
}

func SendCustomErrorResponse(errMessage string) interface{} {
	return map[string]interface{}{
		"status":  constants.API_FAILED_STATUS,
		"message": errMessage,
	}
}
