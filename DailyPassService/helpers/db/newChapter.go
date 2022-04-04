package db

import (
	"encoding/json"
	"fmt"

	structss "github.com/shivamsouravjha/Micro-Game/DailyPassService/services"
	structs "github.com/shivamsouravjha/Micro-Game/DailyPassService/struct"
)

func NewChapter(seriesChapter []byte) {
	var SeriesChapter structs.UserContent
	json.Unmarshal(seriesChapter, &SeriesChapter)
	fmt.Println(SeriesChapter)
	var SeriesDetails []structs.NewChapterContent
	sqlString := fmt.Sprintf("SELECT userId,unlockedChapters FROM `dailypass`")
	structss.Dbmap.Select(&SeriesDetails, sqlString)
	for _, series := range SeriesDetails {
		fmt.Println("series")
		var content map[string]interface{}
		json.Unmarshal([]byte(series.UnlockedChapters), &content)
		content[SeriesChapter.Chapter] = SeriesChapter.ChapterIndex
		contentDB, _ := json.Marshal(content)
		// series.UnlockedChapters[SeriesChapter.Chapter] = SeriesChapter.ChapterIndex
		structss.Dbmap.Exec("UPDATE dailypass SET unlockedChapters = ? where userId = ?", contentDB, series.UserId)
		fmt.Println(SeriesChapter.ChapterIndex)
	}
	return
}
