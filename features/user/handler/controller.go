package handler

import (
	"net/http"
	"strconv"
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
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad request, Incorrect email or password", nil, nil))
		}

		resp, token, err := uh.service.Login(RequestToCore(request))
		if err != nil {
			if strings.Contains(err.Error(), "empty") {
				c.Logger().Error("email and password cannot be empty")
				return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad request, Incorrect email or password", nil, nil))
			}
			if strings.Contains(err.Error(), "invalid") {
				c.Logger().Error("invalid email and password")
				return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad request, Incorrect email or password", nil, nil))
			}
			if strings.Contains(err.Error(), "match") {
				c.Logger().Error("password does not match")
				return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad request, Incorrect email or password", nil, nil))
			}
			if strings.Contains(err.Error(), "token") {
				c.Logger().Error("error while creating jwt token")
				return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "Internal server error", nil, nil))
			}
			if err != nil {
				c.Logger().Error("internal server error")
				return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "Internal server error", nil, nil))
			}
		}

		return c.JSON(http.StatusOK, helper.ResponseFormat(http.StatusOK, "Successful login", loginResponse{
			UserID: resp.UserID, Email: resp.Email, Token: token, Role: resp.Role,
		}, nil))
	}
}

// Register implements user.UserHandler
func (uh *userHandler) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		request := RegisterRequest{}
		_, role, errToken := middlewares.ExtractToken(c)
		if errToken != nil {
			c.Logger().Error("unauthorized access")
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "Missing or Malformed JWT.", nil, nil))
		}
		if role != "admin" {
			c.Logger().Error("unauthorized access")
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "Unauthorized access", nil, nil))
		}

		errBind := c.Bind(&request)
		if errBind != nil {
			c.Logger().Error("error on bind login input")
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad request", nil, nil))
		}

		result, err := uh.service.Register(RequestToCore(request))
		if err != nil {
			if strings.Contains(err.Error(), "empty") {
				return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad request", nil, nil))
			}
			if strings.Contains(err.Error(), "duplicated") {
				return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad request", nil, nil))
			}
			if strings.Contains(err.Error(), "password") {
				return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "Internal server error", nil, nil))
			}
			if err != nil {
				return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "Internal server error", nil, nil))
			}
		}

		resp := register(result)
		return c.JSON(http.StatusCreated, helper.ResponseFormat(http.StatusCreated, "Successfully created an account.", resp, nil))
	}
}

// SearchUser implements user.UserHandler
func (uh *userHandler) SearchUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		keyword := c.QueryParam("keyword")
		limitStr := c.QueryParam("limit")
		offsetStr := c.QueryParam("offset")

		limit := 5
		if limitStr != "" {
			limitInt, err := strconv.Atoi(limitStr)
			if err != nil {
				c.Logger().Errorf("limit is not a number: %s", limitStr)
				return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad Request", nil, nil))
			}
			limit = limitInt
		}

		offset := 0
		if offsetStr != "" {
			offsetInt, err := strconv.Atoi(offsetStr)
			if err != nil {
				c.Logger().Errorf("offset is not a number: %s", offsetStr)
				return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad Request", nil, nil))
			}
			offset = offsetInt
		}

		users, count, err := uh.service.SearchUser(keyword, limit, offset)
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				return c.JSON(http.StatusNotFound, helper.ResponseFormat(http.StatusNotFound, "The requested resource was not found", nil, nil))
			}
			if err != nil {
				return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "Internal server error", nil, nil))
			}
		}

		result := make([]searchUserResponse, len(users))
		for i, user := range users {
			result[i] = searchUser(user)
			result[i].No = uint(i + 1)
		}

		pagination := helper.Paginate(limit, offset, int(count))
		return c.JSON(http.StatusOK, helper.ResponseFormat(http.StatusOK, "Successfully operation", result, pagination))
	}
}

// ProfileUser implements user.UserHandler
func (uh *userHandler) ProfileUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		userId, _, err := middlewares.ExtractToken(c)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "Missing or Malformed JWT.", nil, nil))
		}

		user, err := uh.service.ProfileUser(userId)
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				return c.JSON(http.StatusNotFound, helper.ResponseFormat(http.StatusNotFound, "The requested resource was not found", nil, nil))
			}
			return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "Internal server error", nil, nil))
		}

		resp := profileUser(user)
		return c.JSON(http.StatusOK, helper.ResponseFormat(http.StatusOK, "Successfully operation.", resp, nil))
	}
}
