package pagination

import "github.com/itmrchow/microservice-common/gorm/sort"

type PageReqInfo struct {
	Page  int64            `json:"page" db:"page"`   // 頁碼
	Limit int64            `json:"size" db:"size"`   // 每頁數量
	Sort  []sort.SortOrder `json:"sort" db:"sort"`   // 排序
	Index int64            `json:"index" db:"index"` // 查詢索引
}

type PageRespInfo struct {
	Page       int64
	Limit      int64
	Total      int64
	TotalPages int64
}
