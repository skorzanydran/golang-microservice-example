package user

import (
	"account/internal/group"
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
	groupService      group.Service
	permissionService permission.Service
}

func (r *Controller) Init(G *gin.RouterGroup) {
	G.POST("", r.Create)
	G.GET("", r.FindPage)
	G.GET("/:id", r.FindOne)
	G.PUT("/:id", r.Update)
	G.PUT("/:id/groups/:groupId", r.AddGroup)
	G.DELETE("/:id/groups/:groupId", r.RemoveGroup)
	G.PUT("/:id/permissions/:permissionId", r.AddPermission)
	G.DELETE("/:id/permissions/:permissionId", r.RemovePermission)
}

func (r *Controller) Create(C *gin.Context) {
	validator := CreateRequestValidator{}
	if err := validator.Bind(C); err != nil {
		C.JSON(http.StatusBadRequest, util.NewValidatorError(err))
		return
	}
	err := validator.User.SetPassword(validator.User.Password)
	if err != nil {
		C.JSON(http.StatusBadRequest, util.NewError("user", errors.New("user bad request")))
	}
	user, err := r.Service.Save(validator.User)
	if err != nil {
		C.JSON(http.StatusBadRequest, util.NewError("user", errors.New("user bad request")))
	}

	serializer := Serializer{C, user}
	C.JSON(http.StatusOK, serializer.Response())
}

func (r *Controller) FindPage(C *gin.Context) {
	var users []Model
	page, err := r.Service.FindPage(pageable.InitPageable(util.GetPageable(C.Request.URL.Query())), &users)
	if err != nil {
		C.JSON(http.StatusBadRequest, util.NewError("user", errors.New("user bad request")))
		return
	}

	serializer := ListSerializer{C, users}
	pageSerializer := pageable.PageSerializer{C: C, Page: page}
	C.JSON(http.StatusOK, pageSerializer.Response(&serializer))
}

func (r *Controller) FindOne(C *gin.Context) {
	id, err := uuid.Parse(C.Param("id"))
	if err != nil {
		C.JSON(http.StatusBadRequest, util.NewError("user", errors.New("user bad request")))
		return
	}
	user, err := r.Service.FindOne(id)
	if err != nil {
		C.JSON(http.StatusBadRequest, util.NewError("user", errors.New("user bad request")))
		return
	}

	serializer := Serializer{C, user}
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
	validator.User.Id = id

	user, err := r.Service.Update(validator.User)
	if err != nil {
		C.JSON(http.StatusBadRequest, util.NewError("user", errors.New("user bad request")))
	}

	serializer := Serializer{C, user}
	C.JSON(http.StatusOK, serializer.Response())
}

func (r *Controller) AddGroup(C *gin.Context) {
	id, err := uuid.Parse(C.Param("id"))
	if err != nil {
		C.JSON(http.StatusBadRequest, util.NewError("user", errors.New("not valid UUID")))
		return
	}
	user, err := r.Service.FindOne(id)
	if err != nil {
		C.JSON(http.StatusNotFound, util.NewError("user", errors.New("not Found")))
		return
	}

	groupId, err := uuid.Parse(C.Param("groupId"))
	if err != nil {
		C.JSON(http.StatusBadRequest, util.NewError("group", errors.New("not valid UUID")))
		return
	}
	grp, err := r.groupService.FindOne(groupId)
	if err != nil {
		C.JSON(http.StatusNotFound, util.NewError("group", errors.New("not Found")))
		return
	}

	user, err = r.Service.AddGroup(user, grp)
	if err != nil {
		C.JSON(http.StatusBadRequest, util.NewError("user", errors.New("user bad request")))
		return
	}
	serializer := Serializer{C, user}
	C.JSON(http.StatusOK, serializer.Response())
}

func (r *Controller) RemoveGroup(C *gin.Context) {
	id, err := uuid.Parse(C.Param("id"))
	if err != nil {
		C.JSON(http.StatusBadRequest, util.NewError("user", errors.New("not valid UUID")))
		return
	}
	user, err := r.Service.FindOne(id)
	if err != nil {
		C.JSON(http.StatusNotFound, util.NewError("user", errors.New("not Found")))
		return
	}

	groupId, err := uuid.Parse(C.Param("groupId"))
	if err != nil {
		C.JSON(http.StatusBadRequest, util.NewError("group", errors.New("not valid UUID")))
		return
	}
	grp, err := r.groupService.FindOne(groupId)
	if err != nil {
		C.JSON(http.StatusNotFound, util.NewError("group", errors.New("not Found")))
		return
	}

	user, err = r.Service.RemoveGroup(user, grp)
	if err != nil {
		C.JSON(http.StatusBadRequest, util.NewError("user", errors.New("user bad request")))
		return
	}
	serializer := Serializer{C, user}
	C.JSON(http.StatusOK, serializer.Response())
}

func (r *Controller) AddPermission(C *gin.Context) {
	id, err := uuid.Parse(C.Param("id"))
	if err != nil {
		C.JSON(http.StatusBadRequest, util.NewError("user", errors.New("not valid UUID")))
		return
	}
	user, err := r.Service.FindOne(id)
	if err != nil {
		C.JSON(http.StatusNotFound, util.NewError("user", errors.New("not Found")))
		return
	}

	permissionId, err := uuid.Parse(C.Param("permissionId"))
	if err != nil {
		C.JSON(http.StatusBadRequest, util.NewError("permission", errors.New("not valid UUID")))
		return
	}
	perm, err := r.permissionService.FindOne(permissionId)
	if err != nil {
		C.JSON(http.StatusNotFound, util.NewError("permission", errors.New("not Found")))
		return
	}

	user, err = r.Service.AddPermission(user, perm)
	if err != nil {
		C.JSON(http.StatusBadRequest, util.NewError("user", errors.New("user bad request")))
		return
	}
	serializer := Serializer{C, user}
	C.JSON(http.StatusOK, serializer.Response())
}

func (r *Controller) RemovePermission(C *gin.Context) {
	id, err := uuid.Parse(C.Param("id"))
	if err != nil {
		C.JSON(http.StatusBadRequest, util.NewError("user", errors.New("not valid UUID")))
		return
	}
	user, err := r.Service.FindOne(id)
	if err != nil {
		C.JSON(http.StatusNotFound, util.NewError("user", errors.New("not Found")))
		return
	}

	permissionId, err := uuid.Parse(C.Param("permissionId"))
	if err != nil {
		C.JSON(http.StatusBadRequest, util.NewError("permission", errors.New("not valid UUID")))
		return
	}
	perm, err := r.permissionService.FindOne(permissionId)
	if err != nil {
		C.JSON(http.StatusNotFound, util.NewError("permission", errors.New("not Found")))
		return
	}

	user, err = r.Service.RemovePermission(user, perm)
	if err != nil {
		C.JSON(http.StatusBadRequest, util.NewError("user", errors.New("user bad request")))
		return
	}
	serializer := Serializer{C, user}
	C.JSON(http.StatusOK, serializer.Response())
}
