package db

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	structss "github.com/shivamsouravjha/Micro-Game/ContentService/services"
	structs "github.com/shivamsouravjha/Micro-Game/ContentService/struct"
	requestStruct "github.com/shivamsouravjha/Micro-Game/ContentService/struct/request"
)

func GetContentDAO(ctx context.Context, getContent *requestStruct.GetContent) (*[]structs.ContentDetails, error) {
	var ContentDetails []structs.ContentDetails
	url := "http://localhost:4000/api/v0/getUnlockedContent/" + getContent.UserID + "/" + getContent.SeriesID
	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	bodyBytes, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	var chapterDetails map[string][]string
	err = json.Unmarshal(bodyBytes, &chapterDetails)
	chapterID := chapterDetails["userData"]
	for _, chapter := range chapterID {
		var contentDetail structs.ContentDetails
		sqlString := fmt.Sprintf("SELECT title,story,seriesId FROM `content` WHERE `chapterId` =  \"%v\" ", chapter)
		err = structss.Dbmap.SelectOne(&contentDetail, sqlString)
		if err != nil {
			return nil, err
		}
		ContentDetails = append(ContentDetails, contentDetail)
	}

	if err != nil {
		return nil, err
	}
	return &ContentDetails, nil
}
