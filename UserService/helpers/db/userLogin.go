package db

import (
	"context"
	"fmt"

	structss "github.com/shivamsouravjha/Micro-Game/UserService/services"
	requestStruct "github.com/shivamsouravjha/Micro-Game/UserService/struct/request"
	"github.com/shivamsouravjha/Micro-Game/UserService/utils"
	"golang.org/x/crypto/bcrypt"
)

func UserLoginDAO(ctx context.Context, UserLogin *requestStruct.UserLogin) (string, error) {
	var userPassword string
	sqlString := fmt.Sprintf("select password from user where penname = \"%v\" ", UserLogin.PenName)
	err := structss.Dbmap.SelectOne(&userPassword, sqlString)
	if err != nil {
		fmt.Println(err.Error())
		return err.Error(), err
	}
	err = bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(UserLogin.Password))

	if err != nil {
		return err.Error(), err
	}
	token, err := utils.GenerateToken()

	if err != nil {
		return "Token creation failed", err
	}

	return token, nil

}
