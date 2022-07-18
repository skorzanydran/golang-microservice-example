package pageable

import (
	"github.com/gin-gonic/gin"
)

type PageSerializer struct {
	C    *gin.Context
	Page Page
}

type PageResponse struct {
	Size          int         `json:"size"`
	Page          int         `json:"page"`
	Sort          string      `json:"sort"`
	TotalPages    int         `json:"total_pages"`
	TotalElements int64       `json:"total_elements"`
	Content       interface{} `json:"content"`
}

type Serializer interface {
	Response() interface{}
}

func (s *PageSerializer) Response(serializer Serializer) PageResponse {
	page := PageResponse{
		Size:          s.Page.Size,
		Page:          s.Page.Page,
		Sort:          s.Page.Sort,
		TotalElements: s.Page.TotalElements,
		TotalPages:    s.Page.TotalPages,
	}
	page.Content = serializer.Response()
	return page
}
