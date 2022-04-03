package db

import (
	"context"
	"encoding/json"
	"fmt"

	structss "github.com/shivamsouravjha/Micro-Game/DailyPassService/services"
	structs "github.com/shivamsouravjha/Micro-Game/DailyPassService/struct"
	requestStruct "github.com/shivamsouravjha/Micro-Game/DailyPassService/struct/request"
)

func GetContentDAO(ctx context.Context, getContentRequest *requestStruct.GetUnlockedContent) (*structs.ChapterDetails, error) {
	var ChapterDetails structs.ChapterDetails
	var socialMediaHandles []byte
	sqlString := fmt.Sprintf("SELECT unlockedChapters  FROM `dailypass` WHERE `userId` =  \"%v\" ", getContentRequest.UserId)
	err := structss.Dbmap.SelectOne(&socialMediaHandles, sqlString)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	err = json.Unmarshal(socialMediaHandles, &ChapterDetails.UnlockedContent)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return &ChapterDetails, nil
}
