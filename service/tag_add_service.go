package service

import (
	"github.com/gin-gonic/gin"
	"record-server/model"
	"record-server/serializer"
	"record-server/session"
)

type TagAddService struct {
	TagName string `form:"tag_name" json:"tag_name" binding:"required"`
}

// 添加标签
func (service *TagAddService) TagAdd(c *gin.Context) serializer.Response {
	user := session.CurrentUser(c)
	if user == nil {
		return serializer.ParamErr("尚未登录", nil)
	}

	var tag model.Tag
	tag.TagName = service.TagName
	tag.UserId = user.ID

	if err := model.DB.Create(&tag).Error; err != nil {
		return serializer.ParamErr("标签添加失败", err)
	}

	return serializer.BuildTagResponse(tag)
}
