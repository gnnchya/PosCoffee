package util

type Filters interface{

}

type Filter struct {

}

func NewFilters() (f *Filter) {
	return &Filter{}
}