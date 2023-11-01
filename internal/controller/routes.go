package controller

import (
	"example.com/dynamicWordpressBuilding/internal/middleware"
	"example.com/dynamicWordpressBuilding/utils"
	"github.com/gin-gonic/gin"
)

func (ctl *Controller) Routes() {
	ctl.Router.GET("/", home)
	ctl.publicRoutes()
	//middleware
	authRouter := ctl.Router.Group("/").Use(middleware.AuthMiddleware(utils.NewTokenMaker()))
	ctl.privateRoutes(authRouter)
}

func (ctl *Controller) publicRoutes() {
	ctl.Router.POST("/user", ctl.CreateUser)
	ctl.Router.POST("/user/login", ctl.LoginUser)
	ctl.Router.POST("/user/reset-password", ctl.ResetPassword)
	// ctl.Router.POST("/user/forget-password", ctl.ForgetPassword)
}

func (ctl *Controller) privateRoutes(authRouter gin.IRoutes) {
	authRouter.GET("/alluser", ctl.GetAllUser)
	authRouter.GET("/user/:id", ctl.GetUser)
	authRouter.DELETE("/user/:id", ctl.DeleteUser)
}
