package service

import (
	"github.com/gin-gonic/gin"
	"record-server/model"
	"record-server/serializer"
	"record-server/session"
	"time"
)

type RecordAddService struct {
	EventTime string `form:"event_time" json:"event_time" binding:"required"`
	Comment   string `form:"comment" json:"comment" binding:"required"`
	Amount    int    `form:"amount" json:"amount" binding:"required"`
	Tags      string `form:"tags" json:"tags" binding:"required"`
}

// 添加记录
func (service *RecordAddService) RecordAdd(c *gin.Context) serializer.Response {
	user := session.CurrentUser(c)
	if user == nil {
		return serializer.ParamErr("尚未登录", nil)
	}

	loc, _ := time.LoadLocation("Asia/Shanghai")
	parsedTime, err := time.ParseInLocation("2006-01-02 15:04:05", service.EventTime, loc)
	if err != nil {
		return serializer.ParamErr("参数错误", err)
	}

	var record model.Record
	record.UserId = user.ID
	record.EventTime = parsedTime
	record.Amount = service.Amount
	record.Comment = service.Comment
	record.Tags = service.Tags

	if err := model.DB.Create(&record).Error; err != nil {
		return serializer.ParamErr("记录添加失败", err)
	}

	return serializer.BuildRecordResponse(record)
}
