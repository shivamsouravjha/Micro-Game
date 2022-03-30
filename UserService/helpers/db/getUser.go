package db

import (
	"context"
	"fmt"

	structss "github.com/shivamsouravjha/Micro-Game/services"
	structs "github.com/shivamsouravjha/Micro-Game/struct"
	requestStruct "github.com/shivamsouravjha/Micro-Game/struct/request"
)

func GetCreator(ctx context.Context, getUserRequest *requestStruct.GetUserDetailsRequest) (*structs.UserDetails, error) {
	var UserDetails structs.UserDetails
	sqlString := fmt.Sprintf("SELECT userId,firstName,lastName,penName,userEmail,bio,number FROM `user` WHERE `penName` =  \"%v\" ", getUserRequest.PenName)
	err := structss.Dbmap.SelectOne(&UserDetails, sqlString)
	if err != nil {
		return nil, err
	}
	return &UserDetails, nil
}
