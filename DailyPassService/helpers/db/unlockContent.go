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
	contentCount := fmt.Sprint(usercontent.UnlockedContent[seriesId])
	contentCountInt, _ := strconv.Atoi(contentCount)
	contentCountInt += 1
	usercontent.UnlockedContent[seriesId] = fmt.Sprint(contentCountInt)
	unlockedContents, _ := json.Marshal(usercontent.UnlockedContent)
	_, err = structss.Dbmap.Exec("UPDATE dailypass SET unlockedChapters = ? where userId = ?", string(unlockedContents), userId)

	if err != nil {
		fmt.Println(err.Error())
	}

}
