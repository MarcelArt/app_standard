package models

import "github.com/morkid/paginate"

type IViewable interface {
	ToView() View
}

type View map[string]string

type PageView struct {
	paginate.Page
	Items   []View   `json:"items"`
	Columns []string `json:"columns"`
}
