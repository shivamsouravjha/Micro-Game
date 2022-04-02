package db

import (
	"context"
	"fmt"

	structss "github.com/shivamsouravjha/Micro-Game/ContentService/services"
	structs "github.com/shivamsouravjha/Micro-Game/ContentService/struct"
	requestStruct "github.com/shivamsouravjha/Micro-Game/ContentService/struct/request"
)

func UploadContent(ctx context.Context, createUser *requestStruct.ContentUpload) {
	go taskUploadContent(createUser.Data)
}

func taskUploadContent(content []structs.ContentDetails) {
	for _, element := range content {
		insertedContent, err := structss.Dbmap.Exec("INSERT INTO content (seriesId,story,title) VALUES(?,?,?)", element.SeriesID, element.Story, element.Title)
		if err != nil {
			fmt.Println(err.Error())
		}
		newContentId, _ := insertedContent.LastInsertId()
		_, err = structss.Dbmap.Exec("UPDATE series set chapters=JSON_ARRAY_APPEND(chapters,'$',?) where seriesId=?", newContentId, string(element.SeriesID))
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}
