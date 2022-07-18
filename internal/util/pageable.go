package util

import (
	"net/url"
	"strconv"
)

func GetPageable(values url.Values) (page int, size int, sort string) {
	return GetPageableWithDefault(values, 0, 10, "Created desc")
}

func GetPageableWithDefault(values url.Values, defaultPage int, defaultSize int, defaultSort string) (page int, size int, sort string) {
	p, err := strconv.Atoi(values.Get("page"))
	if err != nil {
		p = defaultPage
	}
	sz, err := strconv.Atoi(values.Get("size"))
	if err != nil {
		sz = defaultSize
	}
	s := values.Get("sort")
	if s == "" {
		s = defaultSort
	}
	return p, sz, s
}
