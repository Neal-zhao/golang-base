package db

import (
	"fmt"
	"blogger/models"
	"github.com/sirupsen/logrus"
	"strings"
)

const TableName = "category"
func GetCategoryInfoByIds(categoryIds []string) (category []models.Category,err error) {
	if len(categoryIds) <= 0 {
		return
	}

	categoryIdStr := strings.Join(categoryIds,",")
	sql := "SELECT * FROM " + TableName + " WHERE 1 "
	sql = fmt.Sprintf("%s AND (%s)",sql,categoryIdStr)
	err = DB.Select(&category,sql,nil)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"msg":fmt.Sprintf("DB.Select err:%s sql:%s",err,sql),
		})
	}
	return
}