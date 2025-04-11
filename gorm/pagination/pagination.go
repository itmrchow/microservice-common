package pagination

import (
	"math"

	"gorm.io/gorm"
)

func OffsetPaginate[T any](model T, pageReq *PageReqInfo, db *gorm.DB) (*gorm.DB, *PageRespInfo, error) {
	// init page resp
	pageResp := &PageRespInfo{
		Page:  pageReq.Page,
		Limit: pageReq.Limit,
	}

	// total
	var total int64
	if err := db.Model(model).Count(&total).Error; err != nil {
		return nil, nil, err
	}
	pageResp.Total = total

	// total pages
	pageResp.TotalPages = int64(math.Ceil(float64(total) / float64(pageResp.Limit)))

	// offset
	offset := (pageReq.Page - 1) * pageReq.Limit

	// query
	query := db.Offset(int(offset)).Limit(int(pageReq.Limit))

	return query, pageResp, nil
}

// func IndexPaginate[T any](model T, pageReq *PageReqInfo, db *gorm.DB) (*gorm.DB, *PageRespInfo, error) {
// 	// TODO: implement
// 	return nil, nil, nil
// }
