package model

type (
	SearchObjectModel struct {
		SearchString string      `json:"searchString"`
		Paging       PagingModel `json:"paging"`
	}

	SearchObjectSystemParamModel struct {
		Category     string      `json:"category"`
		SearchString string      `json:"searchString"`
		Paging       PagingModel `json:"paging"`
	}
)
