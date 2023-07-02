package model

type PageInfo struct {
	Page  int    `json:"page" form:"page"`   //当前页
	Limit int    `json:"limit" form:"limit"` //页大小
	Sort  string `json:"sort" form:"sort"`   //排序字段
}

// 普通分页，通过limit和offset分页
func Pages[T any](listRes *ListResponse[T], pageInfo PageInfo) (err error) {
	var list []T
	var count int64
	if pageInfo.Sort == "" {
		pageInfo.Sort = "id asc" //默认排序条件
	}
	count = Conn.Select("id").Find(&list).RowsAffected
	limit := pageInfo.Limit
	offset := (pageInfo.Page - 1) * limit
	if err = Conn.Limit(limit).Offset(offset).Order(pageInfo.Sort).Find(&list).Error; err != nil {
		return err
	}
	listRes.Count = count
	listRes.List = list
	return nil
}
