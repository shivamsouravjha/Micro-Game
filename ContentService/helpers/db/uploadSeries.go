package db

import (
	"context"
	"encoding/json"
	"fmt"

	structss "github.com/shivamsouravjha/Micro-Game/ContentService/services"
	requestStruct "github.com/shivamsouravjha/Micro-Game/ContentService/struct/request"
)

func UploadSeries(ctx context.Context, uploadSeries *requestStruct.SeriesUpload) {
	chapters, _ := json.Marshal(uploadSeries.Chapters)
	chapterSeries := string(chapters)
	_, err := structss.Dbmap.Exec("INSERT INTO series (author,chapters,name) VALUES(?,?,?)", uploadSeries.Author, chapterSeries, uploadSeries.Name)
	if err != nil {
		fmt.Println(err.Error())
	}
}
