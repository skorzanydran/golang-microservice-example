package group

import (
	"account/internal/pageable"
	"account/internal/permission"
	"account/internal/util"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

type Controller struct {
	Service           Service
	permissionService permission.Service
}

func (r *Controller) Init(G *gin.RouterGroup) {
	G.POST("", r.Create)
	G.GET("", r.FindPage)
	G.GET("/:id", r.FindOne)
	G.PUT("/:id", r.Update)
	G.PUT("/:id/permissions/:permissionId", r.AddPermission)
	G.DELETE("/:id/permissions/:permissionId", r.RemovePermission)
}

func (r *Controller) Create(C *gin.Context) {
	validator := CreateRequestValidator{}
	if err := validator.Bind(C); err != nil {
		C.JSON(http.StatusBadRequest, util.NewValidatorError(err))
		return
	}
	group, err := r.Service.Save(validator.Group)
	if err != nil {
		C.JSON(http.StatusBadRequest, util.NewError("group", errors.New("group bad request")))
	}

	serializer := Serializer{C, group}
	C.JSON(http.StatusOK, serializer.Response())
}

func (r *Controller) FindPage(C *gin.Context) {
	var groups []Model
	page, err := r.Service.FindPage(pageable.InitPageable(util.GetPageable(C.Request.URL.Query())), &groups)
	if err != nil {
		C.JSON(http.StatusBadRequest, util.NewError("group", errors.New("group bad request")))
		return
	}

	serializer := ListSerializer{C, groups}
	pageSerializer := pageable.PageSerializer{C: C, Page: page}
	C.JSON(http.StatusOK, pageSerializer.Response(&serializer))
}

func (r *Controller) FindOne(C *gin.Context) {
	id, err := uuid.Parse(C.Param("id"))
	if err != nil {
		C.JSON(http.StatusBadRequest, util.NewError("group", errors.New("group bad request")))
		return
	}
	group, err := r.Service.FindOne(id)
	if err != nil {
		C.JSON(http.StatusBadRequest, util.NewError("group", errors.New("group bad request")))
		return
	}

	serializer := Serializer{C, group}
	C.JSON(http.StatusOK, serializer.Response())
}

func (r *Controller) Update(C *gin.Context) {
	id, err := uuid.Parse(C.Param("id"))
	if err != nil {
		C.JSON(http.StatusBadRequest, util.NewError("id", errors.New("not valid UUID")))
		return
	}

	validator := UpdateRequestValidator{}
	if err := validator.Bind(C); err != nil {
		C.JSON(http.StatusBadRequest, util.NewValidatorError(err))
		return
	}
	validator.Group.Id = id

	group, err := r.Service.Update(validator.Group)
	if err != nil {
		C.JSON(http.StatusBadRequest, util.NewError("group", errors.New("group bad request")))
	}

	serializer := Serializer{C, group}
	C.JSON(http.StatusOK, serializer.Response())
}

func (r *Controller) AddPermission(C *gin.Context) {
	id, err := uuid.Parse(C.Param("id"))
	if err != nil {
		C.JSON(http.StatusBadRequest, util.NewError("id", errors.New("not valid UUID")))
		return
	}
	group, err := r.Service.FindOne(id)
	if err != nil {
		C.JSON(http.StatusNotFound, util.NewError("group", errors.New("not Found")))
		return
	}

	permissionId, err := uuid.Parse(C.Param("permissionId"))
	if err != nil {
		C.JSON(http.StatusBadRequest, util.NewError("id", errors.New("not valid UUID")))
		return
	}
	perm, err := r.permissionService.FindOne(permissionId)
	if err != nil {
		C.JSON(http.StatusNotFound, util.NewError("permission", errors.New("not Found")))
		return
	}

	group, err = r.Service.AddPermission(group, perm)
	if err != nil {
		C.JSON(http.StatusBadRequest, util.NewError("group", errors.New("group bad request")))
		return
	}
	serializer := Serializer{C, group}
	C.JSON(http.StatusOK, serializer.Response())
}

func (r *Controller) RemovePermission(C *gin.Context) {
	id, err := uuid.Parse(C.Param("id"))
	if err != nil {
		C.JSON(http.StatusBadRequest, util.NewError("id", errors.New("not valid UUID")))
		return
	}
	group, err := r.Service.FindOne(id)
	if err != nil {
		C.JSON(http.StatusNotFound, util.NewError("group", errors.New("not Found")))
		return
	}

	permissionId, err := uuid.Parse(C.Param("permissionId"))
	if err != nil {
		C.JSON(http.StatusBadRequest, util.NewError("id", errors.New("not valid UUID")))
		return
	}
	perm, err := r.permissionService.FindOne(permissionId)
	if err != nil {
		C.JSON(http.StatusNotFound, util.NewError("permission", errors.New("not Found")))
		return
	}

	group, err = r.Service.RemovePermission(group, perm)
	if err != nil {
		C.JSON(http.StatusBadRequest, util.NewError("group", errors.New("group bad request")))
		return
	}
	serializer := Serializer{C, group}
	C.JSON(http.StatusOK, serializer.Response())
}
