package domain

type Paginator struct {
	Total   int
	CurPage int
	PerPage int
	HasMore bool
}

type PageOption struct {
	Page    int      `form:"page,default=1" validate:"min=0"`
	PerPage int      `form:"per_page,default=15" validate:"min=0"`
	Filters []string `form:"filters"`
	Sorts   []string `form:"sorts"`
}

type SetOpParam struct {
	Filters      []string
	SetFieldName string
	Item         interface{}
}
