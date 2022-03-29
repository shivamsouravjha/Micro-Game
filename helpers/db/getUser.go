package db

import (
	"context"
	"database/sql"
	"fmt"

	db "github.com/shivamsouravjha/Micro-Game/services"
	structs "github.com/shivamsouravjha/Micro-Game/struct"
	requestStruct "github.com/shivamsouravjha/Micro-Game/struct/request"
)

func GetCreator(ctx context.Context, getUserRequest *requestStruct.GetUserDetailsRequest) (*structs.UserDetails, error) {
	var UserDetails structs.UserDetails
	fmt.Println("not from redis")
	var sqlString string
	sqlString = fmt.Sprintf("SELECT userId,firstName,lastName,penName,userEmail,number FROM `creatorSpaceUsers` WHERE `handleName` =  \"%v\" AND `isDeleted`= 'false' ", getUserRequest.PenName)
	query := db.WrapQuery(sqlString)

	rows, err := db.GetClient("reader").QueryContext(ctx, query)

	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	rows.Next()
	{

	}

	return &UserDetails, nil

}
