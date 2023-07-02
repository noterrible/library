package model

type ListResponse[T any] struct {
	Count int64 `json:"count" form:"count"`
	List  []T   `json:"list" form:"list"`
}
