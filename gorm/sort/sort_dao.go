package sort

type SortOrder struct {
	Property  string        `json:"property"`  // 排序欄位
	Direction SortDirection `json:"direction"` // 排序方向(ASC/DESC)
}

type SortDirection string

const (
	SortDirectionAsc  SortDirection = "ASC"
	SortDirectionDesc SortDirection = "DESC"
)
