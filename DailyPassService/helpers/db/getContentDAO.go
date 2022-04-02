package db

import (
	"context"
	"encoding/json"
	"fmt"

	structss "github.com/shivamsouravjha/Micro-Game/DailyPassService/services"
	structs "github.com/shivamsouravjha/Micro-Game/DailyPassService/struct"
	requestStruct "github.com/shivamsouravjha/Micro-Game/DailyPassService/struct/request"
)

func GetContentDAO(ctx context.Context, getContentRequest *requestStruct.GetUnlockedContent) (map[string]string, error) {
	var ChapterDetails structs.ChapterDetails
	var socialMediaHandles []byte
	sqlString := fmt.Sprintf("SELECT unlockedSeries  FROM `dailypass` WHERE `userId` =  \"%v\" ", getContentRequest.UserId)
	err := structss.Dbmap.SelectOne(&socialMediaHandles, sqlString)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(socialMediaHandles, &ChapterDetails)
	if err != nil {
		return nil, err
	}
	return ChapterDetails.ChapterId, nil
}
