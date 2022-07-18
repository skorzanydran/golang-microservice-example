package server

import (
	"account/internal/database"
	"account/internal/env"
	"account/internal/group"
	"account/internal/permission"
	"account/internal/user"
	"fmt"
	"github.com/gin-gonic/gin"
)

type Server struct {
	Port int `yaml:"port"`
}

func Start() {
	env.Init()
	database.Init()

	server := Server{}
	env.Load("server", &server)

	permissionService := permission.Service{}
	permissionService.AutoMigrate()

	groupService := group.Service{}
	groupService.AutoMigrate()

	userService := user.Service{}
	userService.AutoMigrate()

	r := gin.Default()

	Router{r, permission.Controller{Service: permissionService}, group.Controller{Service: groupService}, user.Controller{Service: userService}}.Init()

	err := r.Run(fmt.Sprintf(":%d", server.Port))
	if err != nil {
		panic(err)
	}
}
