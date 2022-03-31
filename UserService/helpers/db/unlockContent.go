package db

import (
	"context"
	"fmt"

	structss "github.com/shivamsouravjha/Micro-Game/UserService/services"
	requestStruct "github.com/shivamsouravjha/Micro-Game/UserService/struct/request"
)

func UnlockContentDAO(ctx context.Context, getUserRequest *requestStruct.UnlockContent) {
	go unlockContentEvent(getUserRequest.SeriesID, getUserRequest.ContentID, getUserRequest.UserId)
}

func unlockContentEvent(seriesId int, content []int, userId int) {
	for _, chapter := range content {
		sqlString := fmt.Sprintf("UPDATE user SET items = JSON_ARRAY_APPEND(@unlockedSeries, \"%v\", \"%v\") where userId = \"%v\" ", seriesId, chapter, userId)
		_, err := structss.Dbmap.Exec(sqlString)
		if err != nil {
			fmt.Println(chapter, err.Error())
		}
	}
}
