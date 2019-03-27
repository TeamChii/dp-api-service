package model

type (
	PagingModel struct {
		PageNo      int    `json:"pageNo"`
		PageSize    int    `json:"pageSize"`
		TotalRecord int    `json:"totalRecord"`
		TotalPage   int    `json:"totalPage"`
		OrderBy     string `json:"orderBy"`
		SortBy      string `json:"sortBy"`
	}
)
