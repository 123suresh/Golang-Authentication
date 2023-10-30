package controller

import (
	"net/http"
	"strconv"

	"example.com/dynamicWordpressBuilding/internal/model"
	"example.com/dynamicWordpressBuilding/utils/response"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (controller *Controller) CreateUser(c *gin.Context) {
	userReq := &model.UserRequest{}
	err := c.ShouldBindJSON(userReq)
	if err != nil {
		logrus.Error("json bind error :: ", err)
		response.ERROR(c, err, http.StatusBadRequest)
	}
	userResponse, code, err := controller.svc.CreateUser(userReq)
	if err != nil {
		response.ERROR(c, err, code)
		return
	}
	response.JSON(c, userResponse, "Success", 0, 0)
}

func (controller *Controller) GetAllUser(c *gin.Context) {
	userResponse, code, err := controller.svc.GetAllUser()
	if err != nil {
		response.ERROR(c, err, code)
		return
	}
	response.JSON(c, userResponse, "Success", 0, 0)
}

func (controller *Controller) GetUser(c *gin.Context) {
	id := c.Param("id")
	uid, err := strconv.Atoi(id)
	if err != nil {
		logrus.Error("Error while converting id from string to number")
		return
	}

	userResponse, code, err := controller.svc.GetUser(uid)
	if err != nil {
		response.ERROR(c, err, code)
		return
	}
	response.JSON(c, userResponse, "Success", 0, 0)
}

func (ctl *Controller) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	uid, err := strconv.Atoi(id)
	if err != nil {
		logrus.Error("Error while converting id from string to number")
		return
	}
	code, err := ctl.svc.DeleteUser(uid)
	if err != nil {
		response.ERROR(c, err, code)
		return
	}
	response.JSON(c, "Successfully deleted user", "Success", 0, 0)
}

func (ctl *Controller) LoginUser(c *gin.Context) {
	userLogin := &model.LoginRequest{}
	err := c.ShouldBindJSON(userLogin)
	if err != nil {
		logrus.Error("json bind error :: ", err)
		response.ERROR(c, err, http.StatusBadRequest)
		return
	}
	token, code, err := ctl.svc.LoginUser(userLogin)
	if err != nil {
		response.ERROR(c, err, code)
		return
	}
	response.JSON(c, token, "Success", 0, 0)
}