package handlers

import (
	"bookshelf/models"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h Handler) CreateUser(c *gin.Context) {
	var body models.CreateUserModel
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{Error: "Please enter all requested data"})
		return
	}
	data := c.Request.Method + c.Request.Host + "/myself" + body.Name + body.Secret
	sign := md5.Sum([]byte(data))
	signStr := hex.EncodeToString(sign[:])
	id := uuid.New()
	err := h.IM.AddUser(id.String(), body)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	user, err := h.IM.GetUser(body.Key)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, models.JSONUserResponse{
		Data:    user,
		Sign:    signStr,
		IsOk:    true,
		Message: "ok",
	})
}

func (h Handler) GetUser(c *gin.Context) {
	user, check := h.Check(c)
	if check {
		c.JSON(http.StatusOK, models.JSONGetUserResponse{
			Data:    user,
			IsOk:    true,
			Message: "ok",
		})
	} else {
		c.JSON(http.StatusUnauthorized, models.JSONErrorResponse{
			Error: "Unauthorized!",
		})
	}
}

func (h Handler) Check(c *gin.Context) (models.GetUserModel, bool) {
	key := c.Request.Header.Get("key")
	sign := c.Request.Header.Get("sign")
	fmt.Println(key)
	user, err := h.IM.GetUser(key)
	if err != nil {
		c.JSON(http.StatusNotFound, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return user, false
	}
	data := "POST" + c.Request.Host + c.Request.URL.Path + user.Name + user.Secret
	check := md5.Sum([]byte(data))
	correctStr := hex.EncodeToString(check[:])
	if user.Key == key && sign == correctStr {
		return user, true
	}
	return user, false
}
