package db

import (
	"context"
	"database/sql"
	"fmt"

	structss "github.com/shivamsouravjha/Micro-Game/services"
	requestStruct "github.com/shivamsouravjha/Micro-Game/struct/request"
)

func CreateUser(ctx context.Context, createUser *requestStruct.CreateUserDetails) error {
	_, err := structss.Dbmap.Exec("INSERT INTO user (firstName,lastName,penName,userEmail,bio,number) VALUES(?,?,?,?,?,?)", createUser.FirstName, createUser.LastName, createUser.PenName, createUser.UserEmail, createUser.Bio, createUser.Number)
	if err != nil && err != sql.ErrNoRows {
		fmt.Println(err)
		return err
	}
	return nil
}
