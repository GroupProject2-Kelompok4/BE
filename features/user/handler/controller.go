package handler

import (
	"net/http"
	"strings"

	"github.com/GroupProject2-Kelompok4/BE/features/user"
	"github.com/GroupProject2-Kelompok4/BE/utils/helper"
	"github.com/GroupProject2-Kelompok4/BE/utils/middlewares"
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
			return c.JSON(helper.ResponseFormat(http.StatusBadRequest, "Bad request, Incorrect email or password", nil))
		}

		resp, token, err := uh.service.Login(RequestToCore(request))
		if err != nil {
			if strings.Contains(err.Error(), "empty") {
				c.Logger().Error("email and password cannot be empty")
				return c.JSON(helper.ResponseFormat(http.StatusBadRequest, "Bad request, Incorrect email or password", nil))
			}
			if strings.Contains(err.Error(), "invalid") {
				c.Logger().Error("invalid email and password")
				return c.JSON(helper.ResponseFormat(http.StatusBadRequest, "Bad request, Incorrect email or password", nil))
			}
			if strings.Contains(err.Error(), "match") {
				c.Logger().Error("password does not match")
				return c.JSON(helper.ResponseFormat(http.StatusBadRequest, "Bad request, Incorrect email or password", nil))
			}
			if strings.Contains(err.Error(), "token") {
				c.Logger().Error("error while creating jwt token")
				return c.JSON(helper.ResponseFormat(http.StatusInternalServerError, "Internal server error", nil))
			}
			if err != nil {
				c.Logger().Error("internal server error")
				return c.JSON(helper.ResponseFormat(http.StatusInternalServerError, "Internal server error", nil))
			}
		}

		return c.JSON(helper.ResponseFormat(http.StatusOK, "Successful login", loginResponse{
			UserID: resp.UserID, Email: resp.Email, Token: token, Role: resp.Role,
		}))
	}
}

// Register implements user.UserHandler
func (uh *userHandler) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		request := RegisterRequest{}
		_, role, errToken := middlewares.ExtractToken(c)
		if errToken != nil {
			c.Logger().Error("unauthorized access")
			return c.JSON(helper.ResponseFormat(http.StatusUnauthorized, "Missing or Malformed JWT.", nil))
		}
		if role != "admin" {
			c.Logger().Error("unauthorized access")
			return c.JSON(helper.ResponseFormat(http.StatusUnauthorized, "Unauthorized access", nil))
		}

		errBind := c.Bind(&request)
		if errBind != nil {
			c.Logger().Error("error on bind login input")
			return c.JSON(helper.ResponseFormat(http.StatusBadRequest, "Bad request", nil))
		}

		result, err := uh.service.Register(RequestToCore(request))
		if err != nil {
			if strings.Contains(err.Error(), "empty") {
				return c.JSON(helper.ResponseFormat(http.StatusBadRequest, "Bad request", nil))
			}
			if strings.Contains(err.Error(), "duplicated") {
				return c.JSON(helper.ResponseFormat(http.StatusBadRequest, "Bad request", nil))
			}
			if strings.Contains(err.Error(), "password") {
				return c.JSON(helper.ResponseFormat(http.StatusInternalServerError, "Internal server error", nil))
			}
			if err != nil {
				return c.JSON(helper.ResponseFormat(http.StatusInternalServerError, "Internal server error", nil))
			}
		}

		resp := register(result)
		return c.JSON(helper.ResponseFormat(http.StatusCreated, "Successfully created an account.", resp))
	}
}
