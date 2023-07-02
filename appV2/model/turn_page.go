package model

type TurnPageInfo struct {
	Id               int64  `json:"id" form:"id"`                             //游标，上一页传当前页的第一个id，下一页传入当前页的最后一个
	Page             int    `json:"page" form:"page"`                         //当前页
	Limit            int    `json:"limit" form:"limit"`                       //页大小
	PaginationMethod string `json:"paginationMethod" form:"paginationMethod"` //翻页方式
	Sort             string `json:"sort" form:"sort"`                         //排序方式
}

// 游标分页，通过id和limit分页
func TurnPages[T any](model T, pageInfo TurnPageInfo) (listRes ListResponse[T], err error) {
	var list []T
	var count int64
	if pageInfo.Sort == "" {
		pageInfo.Sort = "id ASC" //升序排序
	}
	count = Conn.Select("id").Find(&list).RowsAffected
	if pageInfo.PaginationMethod == "next" { //下一页
		if err = Conn.Where("id>?", pageInfo.Id).Limit(pageInfo.Limit).Order(pageInfo.Sort).Find(&list).Error; err != nil {
			return listRes, err
		}
	} else if pageInfo.PaginationMethod == "pre" { //上一页
		if err = Conn.Where("id<?", pageInfo.Id).Limit(pageInfo.Limit).Order(pageInfo.Sort).Find(&list).Error; err != nil {
			return listRes, err
		}
	} else { //默认查找第一页10条
		if err = Conn.Where("id>?", 0).Limit(10).Order(pageInfo.Sort + " ASC").Find(&list).Error; err != nil {
			return listRes, err
		}
	}
	listRes.Count = count
	listRes.List = list
	return listRes, nil
}
