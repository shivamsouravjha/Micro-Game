package db

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/shivamsouravjha/Micro-Game/ContentService/config"
	structss "github.com/shivamsouravjha/Micro-Game/ContentService/services"
	structs "github.com/shivamsouravjha/Micro-Game/ContentService/struct"
	requestStruct "github.com/shivamsouravjha/Micro-Game/ContentService/struct/request"
)

func GetContentDAO(ctx context.Context, getContent *requestStruct.GetContent) (*[]structs.ContentDetails, error) {
	var ContentDetails []structs.ContentDetails
	url := config.Get().DAILYPASSSERVICE + "/api/v0/dailypass/getUnlockedContent/" + getContent.UserID
	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	bodyBytes, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	var chapterDetails map[string]map[string]string
	err = json.Unmarshal(bodyBytes, &chapterDetails)
	chapterID, _ := strconv.Atoi(chapterDetails["userData"][getContent.SeriesID])
	fmt.Println(chapterID)
	for i := 1; i <= chapterID; i++ {
		var contentDetails structs.ContentDetails
		sqlString := fmt.Sprintf("SELECT title,story,seriesId FROM `content` WHERE `chapterId` =  \"%v\" ", i)
		err = structss.Dbmap.SelectOne(&contentDetails, sqlString)
		ContentDetails = append(ContentDetails, contentDetails)
	}
	if err != nil {
		return nil, err
	}
	return &ContentDetails, nil
}
