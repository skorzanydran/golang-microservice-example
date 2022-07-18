package server

import (
	"account/internal/group"
	"account/internal/permission"
	"account/internal/user"
	"github.com/gin-gonic/gin"
)

type Router struct {
	E                *gin.Engine
	permissionRouter permission.Controller
	groupRouter      group.Controller
	userRouter       user.Controller
}

func (r Router) Init() {
	g := r.E.Group("/api/v1/account")

	permissions := g.Group("/permissions")
	r.permissionRouter.Init(permissions)

	groups := g.Group("/groups")
	r.groupRouter.Init(groups)

	users := g.Group("/users")
	r.userRouter.Init(users)
}
