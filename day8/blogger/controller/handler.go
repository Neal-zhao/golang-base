package controller

import (
	"blogger/logic"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func IndexHandler(c *gin.Context)  {
	pageSize := c.DefaultQuery("pageSize","10")
	pageNum := c.DefaultQuery("pageNum","1")
	pageSizeInt,_ := strconv.Atoi(pageSize)
	pageNumInt,_ := strconv.Atoi(pageNum)
	logic.GetArticleRecordList(pageNumInt,pageSizeInt)
	c.JSON(http.StatusOK,gin.H{
		"msg":"ok",
	})
}
