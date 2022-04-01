package db

import (
	"context"
	"encoding/json"
	"fmt"

	structss "github.com/shivamsouravjha/Micro-Game/UserService/services"
	requestStruct "github.com/shivamsouravjha/Micro-Game/UserService/struct/request"
)

func UnlockContentDAO(ctx context.Context, getUserRequest *requestStruct.UnlockContent) {
	go unlockContentEvent(ctx, getUserRequest.SeriesID, getUserRequest.ContentID, getUserRequest.UserId)
}

func unlockContentEvent(ctx context.Context, seriesId string, content []string, userId string) {
	for _, chapter := range content {
		getContent := requestStruct.GetUnlockedContent{
			UserId:   userId,
			SeriesId: seriesId,
		}
		usercontent, err := GetContentDAO(ctx, &getContent)
		if err != nil {
			fmt.Println(chapter, err.Error())
		}
		usercontent[seriesId] = append(usercontent[seriesId], chapter)
		seriesContentString, _ := json.Marshal(usercontent)
		userContent := string((seriesContentString[:]))
		sqlString := fmt.Sprintf("UPDATE user SET unlockedSeries = '%v' where userId = \"%v\" ", userContent, userId)
		_, err = structss.Dbmap.Exec(sqlString)
		if err != nil {
			fmt.Println(chapter, err.Error())
		}
	}
}
