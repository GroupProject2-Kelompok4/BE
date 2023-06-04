package handler

import (
	"net/http"
	"strings"

	"github.com/GroupProject2-Kelompok4/BE/features/user"
	"github.com/GroupProject2-Kelompok4/BE/utils/helper"
	"github.com/labstack/echo/v4"
)

type userHandler struct {
	service user.UserService
}

func New(us user.UserService) user.UserHandler {
	return &userHandler{
		service: us,
	}
}

// Login implements user.UserHandler
func (uh *userHandler) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		request := LoginRequest{}
		errBind := c.Bind(&request)
		if errBind != nil {
			c.Logger().Error("error on bind login input")
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad request, Incorrect email or password", nil))
		}

		resp, token, err := uh.service.Login(RequestToCore(request))
		if err != nil {
			if strings.Contains(err.Error(), "empty") {
				c.Logger().Error("email and password cannot be empty")
				return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad request, Incorrect email or password", nil))
			}
			if strings.Contains(err.Error(), "invalid") {
				c.Logger().Error("invalid email and password")
				return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad request, Incorrect email or password", nil))
			}
			if strings.Contains(err.Error(), "match") {
				c.Logger().Error("password does not match")
				return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad request, Incorrect email or password", nil))
			}
			if strings.Contains(err.Error(), "token") {
				c.Logger().Error("error while creating jwt token")
				return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusInternalServerError, "Internal server error", nil))
			}
			if err != nil {
				c.Logger().Error("internal server error")
				return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "Internal server error", nil))
			}
		}

		return c.JSON(http.StatusOK, helper.ResponseFormat(http.StatusOK, "Successful login", LoginResponse{
			UserID: resp.UserID, Email: resp.Email, Token: token, Role: resp.Role,
		}))
	}
}
