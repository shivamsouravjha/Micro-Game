package db

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/shivamsouravjha/Micro-Game/ContentService/helpers/rabbitMQ"
	structss "github.com/shivamsouravjha/Micro-Game/ContentService/services"
	structs "github.com/shivamsouravjha/Micro-Game/ContentService/struct"
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
	newSeries := structs.UserContent{}
	if len(uploadSeries.Chapters) > 4 {
		newSeries.Chapter = fmt.Sprint(insertedSeriesId)
		newSeries.ChapterIndex = fmt.Sprint(4)
	} else {
		newSeries.Chapter = fmt.Sprint(insertedSeriesId)
		newSeries.ChapterIndex = fmt.Sprint(len(uploadSeries.Chapters))
	}

	insertedSeriesIdPush, _ := json.Marshal(newSeries)
	rabbitMQ.RunPublish("SeriesContent", string(insertedSeriesIdPush))
}
