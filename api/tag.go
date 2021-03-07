package api

import (
	"github.com/gin-gonic/gin"
	"record-server/service"
)

// @Title 标签列表
// @Author L. Jiang <l.jiang.1024@gmail.com>
func TagList(c *gin.Context) {
	var tagListService service.TagListService
	res := tagListService.TagList(c)
	c.JSON(0, res)
	return
}

// @Title 添加标签
// @Author L. Jiang <l.jiang.1024@gmail.com>
func TagAdd(c *gin.Context) {
	var tagAddService service.TagAddService
	if err := c.ShouldBind(&tagAddService); err != nil {
		c.JSON(500, ErrorResponse(err))
		return
	}

	res := tagAddService.TagAdd(c)
	c.JSON(0, res)
	return
}
