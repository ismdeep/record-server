package service

import (
	"github.com/gin-gonic/gin"
	"record-server/model"
	"record-server/serializer"
)

type RecordDetailService struct {
}

func (service *RecordDetailService) RecordDetail(c *gin.Context) serializer.Response {
	id := c.Param("id")
	var record model.Record
	if err := model.DB.Select("id,event_time,amount,comment,tags").Where("id=?", id).First(&record).Error; err != nil {
		return serializer.DBErr("查询失败", err)
	}

	return serializer.BuildRecordResponse(record)
}
