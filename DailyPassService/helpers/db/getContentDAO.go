package db

import (
	"encoding/json"
	"fmt"

	structss "github.com/shivamsouravjha/Micro-Game/DailyPassService/services"
	structs "github.com/shivamsouravjha/Micro-Game/DailyPassService/struct"
	requestStruct "github.com/shivamsouravjha/Micro-Game/DailyPassService/struct/request"
)

func GetContentDAO(getContentRequest *requestStruct.GetUnlockedContent, service bool) (*structs.ChapterDetails, error) {
	var ChapterDetails structs.ChapterDetails
	var socialMediaHandles []byte
	sqlString := fmt.Sprintf("SELECT unlockedChapters  FROM `dailypass` WHERE `userId` =  \"%v\" ", getContentRequest.UserId)
	err := structss.Dbmap.SelectOne(&socialMediaHandles, sqlString)
	if err != nil {
		fmt.Println(err.Error(), sqlString)
		return nil, err
	}
	err = json.Unmarshal(socialMediaHandles, &ChapterDetails.UnlockedContent)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	if service {
		ChapterDetailsQueue, _ := json.Marshal(ChapterDetails)
		fmt.Println(ChapterDetailsQueue)
		// rabbitMQ.RunPublish("ReturnContent", string(ChapterDetailsQueue))
	}
	fmt.Println("asd")
	return &ChapterDetails, nil
}
