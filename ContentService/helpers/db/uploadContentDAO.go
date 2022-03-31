package db

import (
	"context"
	"fmt"

	structss "github.com/shivamsouravjha/Micro-Game/ContentService/services"
	structs "github.com/shivamsouravjha/Micro-Game/ContentService/struct"
	requestStruct "github.com/shivamsouravjha/Micro-Game/ContentService/struct/request"
)

func UploadContent(ctx context.Context, createUser *requestStruct.ContentUpload) {
	go taskUpload(createUser.Data)
	return
}

func taskUpload(content []structs.ContentDetails) {
	for _, element := range content {
		_, err := structss.Dbmap.Exec("INSERT INTO user (seriesId,story,title) VALUES(?,?,?)", element.SeriesID, element.Story, element.Title)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}