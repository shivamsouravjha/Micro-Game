package db

import (
	"context"
	"fmt"
	"strings"

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
		var chapterList string
		sqlString := fmt.Sprintf("SELECT chapters FROM `series` WHERE `seriesId` =  \"%v\" ", string(element.SeriesID))
		err = structss.Dbmap.SelectOne(&chapterList, sqlString)
		chapterList = strings.ReplaceAll(chapterList, "]", `,"`+fmt.Sprint(newContentId)+`"]`)
		chapterList = strings.ReplaceAll(chapterList, `"`, `\"`)
		sqlString = fmt.Sprintf("UPDATE series set chapters=\"%v\" where `seriesId`= %v", chapterList, string(element.SeriesID))
		fmt.Println(sqlString)
		_, err = structss.Dbmap.Exec(sqlString)

		if err != nil {
			fmt.Println(err.Error())
		}
	}
}
