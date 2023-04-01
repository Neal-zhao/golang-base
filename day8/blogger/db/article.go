package db

import (
	"blogger/models"
	"fmt"
	"github.com/sirupsen/logrus"
)

const tableName = "article"
func GetArticleInfo(pageNum,pageSize int) (articles []models.ArticleInfo,err error) {
	sql := "SELECT * FROM " + tableName + " WHERE 1 "
	sql = fmt.Sprintf("%s LIMIT %d,%d",sql,(pageNum - 1 ) * pageSize,pageSize)
	err = DB.Select(&articles,sql,nil)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"msg":fmt.Sprintf("DB.Select err:%s sql:%s",err,sql),
		})
	}
	return
}
