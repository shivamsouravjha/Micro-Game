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
}

func taskUpload(content []structs.ContentDetails) {
	for _, element := range content {
		ss, err := structss.Dbmap.Exec("INSERT INTO content (seriesId,story,title) VALUES(?,?,?)", element.SeriesID, element.Story, element.Title)
		if err != nil {
			fmt.Println(err.Error(), ss, element)
		}
	}
}
