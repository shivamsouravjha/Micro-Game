package db

import (
	"context"
	"fmt"

	structss "github.com/shivamsouravjha/Micro-Game/UserService/services"
	structs "github.com/shivamsouravjha/Micro-Game/UserService/struct"
	requestStruct "github.com/shivamsouravjha/Micro-Game/UserService/struct/request"
)

func CreateUserDAO(ctx context.Context, createUser *requestStruct.CreateUserDetails) string {
	var alreadyExisingUser []structs.UserDetails

	sqlString := fmt.Sprintf("SELECT userId,firstName,lastName,penName,userEmail,bio,number FROM `user` WHERE `penName` =  \"%v\"", createUser.PenName)
	_, err := structss.Dbmap.Select(&alreadyExisingUser, sqlString)
	if len(alreadyExisingUser) != 0 || err != nil {
		return "Penname already exists"
	}

	sqlString = fmt.Sprintf("SELECT userId,firstName,lastName,penName,userEmail,bio,number FROM `user` WHERE `userEmail` = \"%v\"", createUser.UserEmail)
	_, err = structss.Dbmap.Select(&alreadyExisingUser, sqlString)
	if len(alreadyExisingUser) != 0 || err != nil {
		return "Email taken"
	}

	_, err = structss.Dbmap.Exec("INSERT INTO user (firstName,lastName,penName,userEmail,bio,number) VALUES(?,?,?,?,?,?)", createUser.FirstName, createUser.LastName, createUser.PenName, createUser.UserEmail, createUser.Bio, createUser.Number)
	if err != nil {
		return err.Error()
	}
	return ""
}
