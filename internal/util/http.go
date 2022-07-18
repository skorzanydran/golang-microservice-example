package util

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func Bind(c *gin.Context, obj interface{}) error {
	b := binding.Default(c.Request.Method, c.ContentType())
	return c.MustBindWith(obj, b)
}

type CommonError struct {
	Errors map[string]interface{} `json:"errors"`
}

type ErrorView struct {
	Status int                    `json:"status"`
	Errors map[string]interface{} `json:"errors"`
}

func HttpError(C *gin.Context, status int, commonError CommonError) {
	C.JSON(status, ErrorView{Status: status, Errors: commonError.Errors})
}

func NewError(key string, err error) CommonError {
	res := CommonError{}
	res.Errors = make(map[string]interface{})
	res.Errors[key] = err.Error()
	return res
}

func NewValidatorError(err error) CommonError {
	res := CommonError{}
	res.Errors = make(map[string]interface{})
	errs := err.(validator.ValidationErrors)
	for _, v := range errs {
		// can translate each error one at a time.
		//fmt.Println("gg",v.NameNamespace)
		if v.Param() != "" {
			res.Errors[v.Field()] = fmt.Sprintf("{%v: %v}", v.Tag(), v.Param())
		} else {
			res.Errors[v.Field()] = fmt.Sprintf("{key: %v}", v.Tag())
		}

	}
	return res
}
