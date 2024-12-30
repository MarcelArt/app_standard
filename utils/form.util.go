package utils

type Form struct {
	Key             string
	Display         string
	Default         string
	Type            string
	DropdownHandler string
	Value           string
}

const (
	BUTTON   = "button"
	CHECKBOX = "checkbox"
	COLOR    = "color"
	DATE     = "date"
	EMAIL    = "email"
	FILE     = "file"
	HIDDEN   = "hidden"
	IMAGE    = "image"
	MONTH    = "month"
	NUMBER   = "number"
	PASSWORD = "password"
	RADIO    = "radio"
	RANGE    = "range"
	RESET    = "reset"
	SEARCH   = "search"
	SUBMIT   = "submit"
	TEL      = "tel"
	TEXT     = "text"
	TIME     = "time"
	URL      = "url"
	WEEK     = "week"
	TEXTAREA = "textarea"
)
