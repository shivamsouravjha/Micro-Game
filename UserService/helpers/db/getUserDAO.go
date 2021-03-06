package db

import (
	"context"
	"fmt"

	structss "github.com/shivamsouravjha/Micro-Game/UserService/services"
	structs "github.com/shivamsouravjha/Micro-Game/UserService/struct"
)

func GetCreatorDAO(ctx context.Context, penName string) (*structs.UserDetails, error) {
	var UserDetails structs.UserDetails
	sqlString := fmt.Sprintf("SELECT userId,firstName,lastName,penName,userEmail,bio,number FROM `user` WHERE `penName` =  \"%v\" ", penName)
	err := structss.Dbmap.SelectOne(&UserDetails, sqlString)
	if err != nil {
		return nil, err
	}
	return &UserDetails, nil
}
