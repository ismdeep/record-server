package api

import (
	"github.com/gin-gonic/gin"
	"record-server/service"
)

// @Title 添加打卡记录
// @Author L. Jiang <l.jiang.1024@gmail.com>
func RecordAdd(c *gin.Context) {
	var recordAddService service.RecordAddService
	if err := c.ShouldBind(&recordAddService); err != nil {
		c.JSON(500, ErrorResponse(err))
		return
	}

	res := recordAddService.RecordAdd(c)
	c.JSON(0, res)
	return
}

// @Title 查看所有打卡记录
// @Author L. Jiang <l.jiang.1024@gmail.com>
func RecordList(c *gin.Context) {
	var serv service.RecordListService
	c.JSON(0, serv.RecordList(c))
	return
}

// @Title 查看打开记录详情
// @Author L. Jiang <l.jiang.1024@gmail.com>
func RecordDetail(c *gin.Context) {
	var serv service.RecordDetailService
	c.JSON(0, serv.RecordDetail(c))
	return
}

// @Title 删除打卡记录
// @Author L. Jiang <l.jiang.1024@gmail.com>
func RecordDelete(c *gin.Context) {
	var serv service.RecordDeleteService
	c.JSON(0, serv.RecordDelete(c))
	return
}
