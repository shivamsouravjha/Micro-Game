package db

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/shivamsouravjha/Micro-Game/ContentService/helpers/rabbitMQ"
	structss "github.com/shivamsouravjha/Micro-Game/ContentService/services"
	requestStruct "github.com/shivamsouravjha/Micro-Game/ContentService/struct/request"
)

func UploadSeries(ctx context.Context, uploadSeries *requestStruct.SeriesUpload) {
	chapters, _ := json.Marshal(uploadSeries.Chapters)
	chapterSeries := string(chapters)
	insertedSeries, err := structss.Dbmap.Exec("INSERT INTO series (author,chapters,name) VALUES(?,?,?)", uploadSeries.Author, chapterSeries, uploadSeries.Name)
	if err != nil {
		fmt.Println(err.Error())
	}

	insertedSeriesId, _ := insertedSeries.LastInsertId()
	newSeries := make(map[int]interface{})
	if len(uploadSeries.Chapters) > 4 {
		newSeries[int(insertedSeriesId)] = 4
	} else {
		newSeries[int(insertedSeriesId)] = len(uploadSeries.Chapters)
	}

	insertedSeriesIdPush := fmt.Sprintln(newSeries)
	rabbitMQ.RunPublish("SeriesContent", insertedSeriesIdPush)
}
