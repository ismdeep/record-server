package service

import (
	"github.com/gin-gonic/gin"
	"log"
	"record-server/model"
	"record-server/serializer"
)

type RecordDeleteService struct {
}

func (service *RecordDeleteService) RecordDelete(c *gin.Context) serializer.Response {
	id := c.Param("id")
	log.Println(id)

	var record model.Record

	if err := model.DB.Where("id = ?", id).Delete(&record).Error; err != nil {
		return serializer.ParamErr("删除错误", err)
	}

	return serializer.Response{
		Code:  0,
		Data:  nil,
		Msg:   "",
		Error: "",
	}
}
