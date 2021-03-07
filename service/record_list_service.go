package service

import (
	"github.com/gin-gonic/gin"
	"record-server/model"
	"record-server/serializer"
	"record-server/session"
)

type RecordListService struct {
}

func (service *RecordListService) RecordList(c *gin.Context) serializer.Response {
	user := session.CurrentUser(c)
	if user == nil {
		return serializer.ParamErr("尚未登录", nil)
	}

	var RecordList []model.Record

	if err := model.DB.Select("id,event_time,amount,comment,tags,created_at").Where("user_id = ?", user.ID).Find(&RecordList).Error; err != nil {
		return serializer.ParamErr("查询错误", err)
	}

	return serializer.BuildRecordListResponse(RecordList)
}
