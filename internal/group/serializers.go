package group

import (
	"account/internal/permission"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"time"
)

type Serializer struct {
	C *gin.Context
	Model
}

type Response struct {
	Id          uuid.UUID             `json:"id"`
	Name        string                `json:"name"`
	Created     time.Time             `json:"created"`
	Updated     time.Time             `json:"updated"`
	Permissions []permission.Response `json:"permissions"`
}

func (s *Serializer) Response() Response {
	permissions := []permission.Response{}
	for _, perm := range s.Permissions {
		serializer := permission.Serializer{C: s.C, Model: perm}
		permissions = append(permissions, serializer.Response())
	}

	response := Response{
		Id:          s.Id,
		Name:        s.Name,
		Created:     s.Created,
		Updated:     s.Updated,
		Permissions: permissions,
	}
	return response
}

type ListSerializer struct {
	C      *gin.Context
	Groups []Model
}

func (s *ListSerializer) Response() interface{} {
	response := []Response{}
	for _, group := range s.Groups {
		serializer := Serializer{s.C, group}
		response = append(response, serializer.Response())
	}
	return response
}
