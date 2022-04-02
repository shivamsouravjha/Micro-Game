package db

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	structss "github.com/shivamsouravjha/Micro-Game/DailyPassService/services"
	requestStruct "github.com/shivamsouravjha/Micro-Game/DailyPassService/struct/request"
)

func UnlockContentDAO(ctx context.Context, getUserRequest *requestStruct.UnlockContent) {
	go unlockContentEvent(ctx, getUserRequest.SeriesID, getUserRequest.UserId)
}

func unlockContentEvent(ctx context.Context, seriesId string, userId string) {
	getContent := requestStruct.GetUnlockedContent{
		UserId: userId,
	}
	usercontent, err := GetContentDAO(ctx, &getContent)
	if err != nil {
		fmt.Println(err.Error())
	}
	contentCount, _ := strconv.Atoi(usercontent[seriesId])
	contentCount += 1
	usercontent[seriesId] = fmt.Sprint(contentCount)
	seriesContentString, _ := json.Marshal(usercontent)
	userContent := string((seriesContentString[:]))
	sqlString := fmt.Sprintf("UPDATE dailypass SET unlockedChapter = JSON_SET(unlockedChapter,'$',\"%v\",) where userId = \"%v\" ", userContent, userId)
	_, err = structss.Dbmap.Exec(sqlString)
	if err != nil {
		fmt.Println(err.Error())
	}

}
