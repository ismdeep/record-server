package service

import (
	"github.com/gin-gonic/gin"
	"record-server/model"
	"record-server/serializer"
	"record-server/session"
)

type TagListService struct {
}

// 添加标签
func (service *TagListService) TagList(c *gin.Context) serializer.Response {
	user := session.CurrentUser(c)
	if user == nil {
		return serializer.ParamErr("尚未登录", nil)
	}

	var TagList []model.Tag

	if err := model.DB.Select("id,tag_name,created_at").Where("user_id = ?", user.ID).Find(&TagList).Error; err != nil {
		return serializer.ParamErr("查询错误", err)
	}

	return serializer.BuildTagListResponse(TagList)
}
