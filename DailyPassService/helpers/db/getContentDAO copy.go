package db

import (
	"encoding/json"
	"fmt"

	structss "github.com/shivamsouravjha/Micro-Game/DailyPassService/services"
	structs "github.com/shivamsouravjha/Micro-Game/DailyPassService/struct"
)

func newUserChapter(SeriesChapter string) (*structs.ChapterDetails, error) {
	var ChapterDetails structs.ChapterDetails
	var socialMediaHandles []byte
	sqlString := fmt.Sprintf("SELECT unlockedChapters  FROM `dailypass` WHERE `userId` =  \"%v\" ")
	err := structss.Dbmap.SelectOne(&socialMediaHandles, sqlString)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	err = json.Unmarshal(socialMediaHandles, &ChapterDetails.UnlockedContent)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return &ChapterDetails, nil
}
