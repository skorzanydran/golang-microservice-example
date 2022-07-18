package user

import (
	"account/internal/util"
	"github.com/gin-gonic/gin"
)

type CreateRequestValidator struct {
	Username string `form:"username" json:"username" binding:"required,min=5,max=50"`
	Password string `form:"password" json:"password" binding:"required,min=6,max=255"`

	User Model `json:"-"`
}

func (v *CreateRequestValidator) Bind(C *gin.Context) error {
	err := util.Bind(C, v)
	if err != nil {
		return err
	}

	v.User.Username = v.Username
	v.User.Password = v.Password

	return nil
}

type UpdateRequestValidator struct {
	Username string `form:"username" json:"username" binding:"required,min=5,max=50"`

	User Model `json:"-"`
}

func (v *UpdateRequestValidator) Bind(C *gin.Context) error {
	err := util.Bind(C, v)
	if err != nil {
		return err
	}

	v.User.Username = v.Username

	return nil
}
