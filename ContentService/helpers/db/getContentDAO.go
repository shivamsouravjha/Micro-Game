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
	//send userid and series id and get array of content
	url := "localhost:4000/api/v0/getUnlockedContent" + getContent.UserID + getContent.SeriesID
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	bodyBytes, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	var chapterId []string
	err = json.Unmarshal(bodyBytes, &chapterId)
	for _, chapter := range chapterId {
		var contentDetail structs.ContentDetails
		sqlString := fmt.Sprintf("SELECT title,story FROM `content` WHERE `chapterId` =  \"%v\" ", chapter)
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
