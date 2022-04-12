package handler

import (
	user "go-api/User"
	"go-api/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	//catch input from user
	//map input from user to struc registeruserinput
	//strct di atas passing sebagai parameter ke service
	input := user.RegisterUserInput{}
	err := c.ShouldBind(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Register account failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		response := helper.APIResponse("Register account failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(newUser, "eyKa78akjdaw17kuiydajkd")
	response := helper.APIResponse("Account has ben registered", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}
