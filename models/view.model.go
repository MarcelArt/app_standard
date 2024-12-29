package models

import "github.com/morkid/paginate"

type View map[string]string

func (v View) Keys() []string {
	var keys []string
	for key := range v {
		keys = append(keys, key)
	}
	return keys
}

type IViewable interface {
	ToView() View
}

type PageView struct {
	paginate.Page
	Items       []View   `json:"items"`
	Columns     []string `json:"columns"`
	Title       string
	Description string
}
