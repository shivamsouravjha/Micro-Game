package db

import (
	"context"
	"fmt"

	structss "github.com/shivamsouravjha/Micro-Game/ContentService/services"
	structs "github.com/shivamsouravjha/Micro-Game/ContentService/struct"
	requestStruct "github.com/shivamsouravjha/Micro-Game/ContentService/struct/request"
)

func GetSeriesDAO(ctx context.Context, getSeries *requestStruct.GetSeries) (*structs.SeriesDetails, error) {
	var SeriesDetails structs.SeriesDetails
	sqlString := fmt.Sprintf("SELECT seriesId,author,chapters,name FROM `series` WHERE `seriesId` =  \"%v\" ", getSeries.SeriesID)
	err := structss.Dbmap.SelectOne(&SeriesDetails, sqlString)
	if err != nil {
		return nil, err
	}
	return &SeriesDetails, nil
}
