package permission

import (
	"account/internal/util"
	"github.com/gin-gonic/gin"
)

type CreateRequestValidator struct {
	Name string `form:"name" json:"name" binding:"required,min=5,max=50"`

	Permission Model `json:"-"`
}

func (v *CreateRequestValidator) Bind(C *gin.Context) error {
	err := util.Bind(C, v)
	if err != nil {
		return err
	}

	v.Permission.Name = v.Name

	return nil
}

type UpdateRequestValidator struct {
	Name string `form:"name" json:"name" binding:"required,min=5,max=50"`

	Permission Model `json:"-"`
}

func (v *UpdateRequestValidator) Bind(C *gin.Context) error {
	err := util.Bind(C, v)
	if err != nil {
		return err
	}

	v.Permission.Name = v.Name

	return nil
}
