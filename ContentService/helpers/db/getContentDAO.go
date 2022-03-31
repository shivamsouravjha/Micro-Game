package db

import (
	"context"
	"fmt"

	structss "github.com/shivamsouravjha/Micro-Game/ContentService/services"
	structs "github.com/shivamsouravjha/Micro-Game/ContentService/struct"
	requestStruct "github.com/shivamsouravjha/Micro-Game/ContentService/struct/request"
)

func GetContentDAO(ctx context.Context, getContent *requestStruct.GetContent) (*structs.ContentDetails, error) {
	var UserDetails structs.ContentDetails
	//send userid and series id and get array of content
	sqlString := fmt.Sprintf("SELECT title,story FROM `content` WHERE `chapterId` =  \"%v\" ", getContent.SeriesID)

	err := structss.Dbmap.SelectOne(&UserDetails, sqlString)
	if err != nil {
		return nil, err
	}
	return &UserDetails, nil
}
