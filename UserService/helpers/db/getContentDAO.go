package db

import (
	"context"
	"fmt"

	structss "github.com/shivamsouravjha/Micro-Game/UserService/services"
	structs "github.com/shivamsouravjha/Micro-Game/UserService/struct"
	requestStruct "github.com/shivamsouravjha/Micro-Game/UserService/struct/request"
)

func GetContentDAO(ctx context.Context, getContentRequest *requestStruct.GetUnlockedContent) ([]int, error) {
	var ChapterDetails structs.UserDetails
	sqlString := fmt.Sprintf("SELECT series FROM `user` WHERE `penName` =  \"%v\" ", getContentRequest.UserId)
	err := structss.Dbmap.SelectOne(&ChapterDetails, sqlString)
	if err != nil {
		return nil, err
	}
	return ChapterDetails.UnlockedSeries[getContentRequest.SeriesId], nil
}
