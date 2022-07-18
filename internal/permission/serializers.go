package permission

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"time"
)

type Serializer struct {
	C *gin.Context
	Model
}

type Response struct {
	Id      uuid.UUID `json:"id"`
	Name    string    `json:"name"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
}

func (s *Serializer) Response() Response {
	response := Response{
		Id:      s.Id,
		Name:    s.Name,
		Created: s.Created,
		Updated: s.Updated,
	}
	return response
}

type ListSerializer struct {
	C           *gin.Context
	Permissions []Model
}

func (s *ListSerializer) Response() interface{} {
	response := []Response{}
	for _, permission := range s.Permissions {
		serializer := Serializer{s.C, permission}
		response = append(response, serializer.Response())
	}
	return response
}
