package db

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	queue "github.com/shivamsouravjha/Micro-Game/ContentService/helpers/rabbitMQ"
	structss "github.com/shivamsouravjha/Micro-Game/ContentService/services"
	structs "github.com/shivamsouravjha/Micro-Game/ContentService/struct"
	requestStruct "github.com/shivamsouravjha/Micro-Game/ContentService/struct/request"
)

func GetContentDAO(ctx context.Context, getContent *requestStruct.GetContent) (*structs.ContentDetails, error) {
	var ContentDetails structs.ContentDetails
	queue.Run("ds")
	url := "http://localhost:8003/api/v0/dailypass/getUnlockedContent/" + getContent.UserID
	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	bodyBytes, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	var chapterDetails map[string]map[string]interface{}
	err = json.Unmarshal(bodyBytes, &chapterDetails)
	chapterID := chapterDetails["userData"][getContent.SeriesID]
	sqlString := fmt.Sprintf("SELECT title,story,seriesId FROM `content` WHERE `chapterId` =  \"%v\" ", chapterID)
	err = structss.Dbmap.SelectOne(&ContentDetails, sqlString)
	if err != nil {
		return nil, err
	}
	return &ContentDetails, nil
}
