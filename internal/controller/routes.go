package controller

func (ctl *Controller) Routes() {
	ctl.Router.GET("/", home)
	ctl.userRoutes()
}

func (ctl *Controller) userRoutes() {
	ctl.Router.POST("/user", ctl.CreateUser)
	ctl.Router.GET("/alluser", ctl.GetAllUser)
	ctl.Router.GET("/user/:id", ctl.GetUser)
	ctl.Router.DELETE("/user/:id", ctl.DeleteUser)
	ctl.Router.POST("/user/login", ctl.LoginUser)
}
