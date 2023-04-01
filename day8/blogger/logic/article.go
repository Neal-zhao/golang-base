package logic

import (
	"blogger/db"
	"blogger/models"
)

func GetArticleRecordList(pageNum,pageSize int)  {
	articleInfo,err := db.GetArticleInfo(pageNum,pageSize)
	if err != nil {
		return
	}

	//var articleById []map[int64]*models.ArticleInfo
	var categoryIds []string
	for _,v := range articleInfo {
		categoryIds = append(categoryIds,string(v.Id))
		//articleById = append(articleById,map[int64]*models.ArticleInfo{
		//	v.Id:&v,
		//})
	}

	categoryInfo,err := db.GetCategoryInfoByIds(categoryIds)
	if err != nil {
		return
	}

	categoryById := make(map[int64]models.Category)
	for _,v := range categoryInfo {
		categoryById[v.Id] = v
	}

	var articleRecord []models.ArticleRecord
	for _,v := range articleInfo {
		categoryIds = append(articleRecord,v)
		//articleById = append(articleById,map[int64]*models.ArticleInfo{
		//	v.Id:&v,
		//})
	}
}