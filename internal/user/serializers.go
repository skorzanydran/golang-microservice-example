package user

import (
	"account/internal/group"
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
	Username    string                `json:"name"`
	Groups      []group.Response      `json:"groups"`
	Permissions []permission.Response `json:"permissions"`
	Created     time.Time             `json:"created"`
	Updated     time.Time             `json:"updated"`
}

func (s *Serializer) Response() Response {
	permissions := []permission.Response{}
	for _, perm := range s.Permissions {
		serializer := permission.Serializer{C: s.C, Model: perm}
		permissions = append(permissions, serializer.Response())
	}
	groups := []group.Response{}
	for _, grp := range s.Groups {
		serializer := group.Serializer{C: s.C, Model: grp}
		groups = append(groups, serializer.Response())
	}

	response := Response{
		Id:          s.Id,
		Username:    s.Username,
		Created:     s.Created,
		Updated:     s.Updated,
		Groups:      groups,
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
	for _, user := range s.Groups {
		serializer := Serializer{s.C, user}
		response = append(response, serializer.Response())
	}
	return response
}
