package result

type PageResult[T any] struct {
	TotalCount int `json:"totalCount"`
	Item       T   `json:"item"`
}
