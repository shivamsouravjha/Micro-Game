package db

import (
	"context"
	"database/sql"
	"fmt"

	structss "github.com/shivamsouravjha/Micro-Game/services"
	structs "github.com/shivamsouravjha/Micro-Game/struct"
	requestStruct "github.com/shivamsouravjha/Micro-Game/struct/request"
)

func GetCreator(ctx context.Context, getUserRequest *requestStruct.GetUserDetailsRequest) (*structs.UserDetails, error) {
	var UserDetails structs.UserDetails
	var sqlString string
	sqlString = fmt.Sprintf("SELECT userId,firstName,lastName,penName,userEmail,bio,number FROM `user` WHERE `penName` =  \"%v\" ", getUserRequest.PenName)
	fmt.Println(sqlString)
	err := structss.Dbmap.SelectOne(&UserDetails, sqlString)
	fmt.Println(UserDetails.PenName)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return &UserDetails, nil
}
