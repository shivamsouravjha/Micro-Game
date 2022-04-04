package rabbitMQ

import (
	"encoding/json"
	"fmt"

	structss "github.com/shivamsouravjha/Micro-Game/ContentService/services"
	structs "github.com/shivamsouravjha/Micro-Game/ContentService/struct"
)

func NewUserContent(userId string) string {
	var SeriesDetails []structs.SeriesDetails
	var UserContentDetails structs.NewUserDetails
	UserContentDetails.UnlockedChapters = make(map[string]interface{})
	UserContentDetails.UserId = userId
	sqlString := fmt.Sprintf("SELECT chapters,seriesId FROM `series`")
	structss.Dbmap.Select(&SeriesDetails, sqlString)
	for _, series := range SeriesDetails {
		if len(series.Chapters) > 4 {
			UserContentDetails.UnlockedChapters[series.SeriesID] = 4
		} else {
			UserContentDetails.UnlockedChapters[series.SeriesID] = len(series.Chapters)
		}
	}
	responseUserDetails, _ := json.Marshal(UserContentDetails)
	return string(responseUserDetails)
}
