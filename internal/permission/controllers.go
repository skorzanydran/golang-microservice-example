package permission

import (
	"account/internal/pageable"
	"account/internal/util"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

type Controller struct {
	Service Service
}

func (r *Controller) Init(G *gin.RouterGroup) {
	G.POST("", r.Create)
	G.GET("", r.FindPage)
	G.GET("/:id", r.FindOne)
	G.PUT("/:id", r.Update)
}

func (r *Controller) Create(C *gin.Context) {
	validator := CreateRequestValidator{}
	if err := validator.Bind(C); err != nil {
		util.HttpError(C, http.StatusBadRequest, util.NewValidatorError(err))
		return
	}
	permission, err := r.Service.Save(validator.Permission)
	if err != nil {
		util.HttpError(C, http.StatusBadRequest, util.NewError("permission", errors.New("permission bad request")))
	}

	serializer := Serializer{C, permission}
	C.JSON(http.StatusOK, serializer.Response())
}

func (r *Controller) Update(C *gin.Context) {
	id, err := uuid.Parse(C.Param("id"))
	if err != nil {
		util.HttpError(C, http.StatusBadRequest, util.NewError("id", errors.New("not valid UUID")))
		return
	}

	validator := UpdateRequestValidator{}
	if err := validator.Bind(C); err != nil {
		util.HttpError(C, http.StatusBadRequest, util.NewValidatorError(err))
		return
	}
	validator.Permission.Id = id

	permission, err := r.Service.Update(validator.Permission)
	if err != nil {
		util.HttpError(C, http.StatusBadRequest, util.NewError("permission", errors.New("permission bad request")))
	}

	serializer := Serializer{C, permission}
	C.JSON(http.StatusOK, serializer.Response())
}

func (r *Controller) FindPage(C *gin.Context) {
	var permissions []Model
	page, err := r.Service.FindPage(pageable.InitPageable(util.GetPageable(C.Request.URL.Query())), &permissions)
	if err != nil {
		util.HttpError(C, http.StatusBadRequest, util.NewError("permission", errors.New("permission bad request")))
		return
	}

	serializer := ListSerializer{C, permissions}
	pageSerializer := pageable.PageSerializer{C: C, Page: page}
	C.JSON(http.StatusOK, pageSerializer.Response(&serializer))
}

func (r *Controller) FindOne(C *gin.Context) {
	id, err := uuid.Parse(C.Param("id"))
	if err != nil {
		util.HttpError(C, http.StatusBadRequest, util.NewError("permission", errors.New("permission bad request")))
		return
	}
	permission, err := r.Service.FindOne(id)
	if err != nil {
		util.HttpError(C, http.StatusBadRequest, util.NewError("permission", errors.New("permission bad request")))
		return
	}

	serializer := Serializer{C, permission}
	C.JSON(http.StatusOK, serializer.Response())
}
