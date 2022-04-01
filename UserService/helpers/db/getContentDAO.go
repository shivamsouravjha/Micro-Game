package db

import (
	"context"
	"encoding/json"
	"fmt"

	structss "github.com/shivamsouravjha/Micro-Game/UserService/services"
	structs "github.com/shivamsouravjha/Micro-Game/UserService/struct"
	requestStruct "github.com/shivamsouravjha/Micro-Game/UserService/struct/request"
)

func GetContentDAO(ctx context.Context, getContentRequest *requestStruct.GetUnlockedContent) (map[string][]interface{}, error) {
	var ChapterDetails structs.UserDetails
	var socialMediaHandles []byte

	sqlString := fmt.Sprintf("SELECT unlockedSeries FROM `user` WHERE `userId` =  \"%v\" ", getContentRequest.UserId)
	err := structss.Dbmap.SelectOne(&socialMediaHandles, sqlString)

	err = json.Unmarshal(socialMediaHandles, &ChapterDetails.UnlockedSeries)
	if err != nil {
		return nil, err
	}
	return ChapterDetails.UnlockedSeries, nil //.// [getContentRequest.SeriesId], nil
}
