package group

import (
	"account/internal/util"
	"github.com/gin-gonic/gin"
)

type CreateRequestValidator struct {
	Name string `form:"name" json:"name" binding:"required,min=5,max=50"`

	Group Model `json:"-"`
}

func (v *CreateRequestValidator) Bind(C *gin.Context) error {
	err := util.Bind(C, v)
	if err != nil {
		return err
	}

	v.Group.Name = v.Name

	return nil
}

type UpdateRequestValidator struct {
	Name string `form:"name" json:"name" binding:"required,min=5,max=50"`

	Group Model `json:"-"`
}

func (v *UpdateRequestValidator) Bind(C *gin.Context) error {
	err := util.Bind(C, v)
	if err != nil {
		return err
	}

	v.Group.Name = v.Name

	return nil
}
