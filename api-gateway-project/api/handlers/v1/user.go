package v1

import (
	"context"
	"encoding/json"
	models "exam_task_4/api-gateway-project/api/handlers/models"
	pbu "exam_task_4/api-gateway-project/genproto/user-service"
	l "exam_task_4/api-gateway-project/pkg/logger"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"google.golang.org/protobuf/encoding/protojson"
)

// CreateUser
// @Summary CreateUser
// @Security ApiKeyAuth
// @Description Api for CreateUser
// @Tags user
// @Accept json
// @Produce json
// @Param User body models.CreateUser true "createUserModel"
// @Success 200 {object} models.RespUser
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/user/create [post]
func (h *handlerV1) CreateUser(c *gin.Context) {

	var (
		body        models.User
		jsonMarshal protojson.MarshalOptions
	)
	jsonMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", l.Error(err))
		return
	}

	id := uuid.New().String()
	body.Id = id
	dataBytes, err := json.Marshal(&body)
	if err != nil {
		log.Println(err)
	}

	h.Writer.ProducerMessage("createUser", dataBytes)

	// ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	// defer cancel()
	// response, err := h.serviceManager.UserService().CreateUser(ctx, &pbu.User{
	// 	FirstName: body.FirstName,
	// 	LastName:  body.LastName,
	// 	Bio:       body.Bio,
	// 	Website:   body.Website,
	// 	Username:  body.UserName,
	// 	Email:     body.Email,
	// 	Password:  body.Password,
	// })
	resp := models.RespUser{
		Id:        body.Id,
		UserName:  body.FirstName,
		Email:     body.LastName,
		Password:  body.Bio,
		FirstName: body.Website,
		LastName:  body.UserName,
		Bio:       body.Email,
		Website:   body.Password,
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to create user", l.Error(err))
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// GetUser
// @Summary GetUser
// @Security ApiKeyAuth
// @Description Api for GetUser
// @Tags user
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Success 200 {object} models.RespUser
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/user/{id} [get]
func (h *handlerV1) GetUser(c *gin.Context) {

	var jsonMarshal protojson.MarshalOptions
	jsonMarshal.UseProtoNames = true

	id := c.Param("id")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.UserService().GetUser(
		ctx, &pbu.UserId{
			UserId: id,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get user", l.Error(err))
		return
	}
	resp := models.RespUser{
		Id:        response.Id,
		FirstName: response.FirstName,
		LastName:  response.LastName,
		Bio:       response.Bio,
		Website:   response.Website,
		UserName:  response.Username,
		Email:     response.Email,
		Password:  response.Password,
	}

	c.JSON(http.StatusOK, resp)
}

// ListUsers gets user by id
// @Summary ListUsers
// @Description Api for getting all users
// @Tags user
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param page query string true "User PAGES"
// @Param limit query string true "User LIMIT"
// @Success 200 {object} models.RespUser
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/getall [GET]
func (h *handlerV1) ListUsers(c *gin.Context) {
	page := c.Query("page")
	limit := c.Query("limit")

	PageNum, err := strconv.Atoi(page)
	if err != nil {
		return
	}

	LimitNum, err := strconv.Atoi(limit)
	if err != nil {
		return
	}

	ctxWithCancel, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.UserService().GetAllUsers(
		ctxWithCancel, &pbu.GetAllUsersRequest{
			Page:  int64(PageNum),
			Limit: int64(LimitNum),
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to list users")
		return
	}
	var respUsers []models.RespUser
	for _, user := range response.Users {
		resp := models.RespUser{
			Id:        user.Id,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Bio:       user.Bio,
			Website:   user.Website,
			UserName:  user.Username,
			Email:     user.Email,
			Password:  user.Password,
		}
		respUsers = append(respUsers, resp)
	}

	c.JSON(http.StatusOK, respUsers)
}

// UpdateUser
// @Summary UpdateUser
// @Security ApiKeyAuth
// @Description Api for UpdateUser
// @Tags user
// @Accept json
// @Produce json
// @Param User body models.User true "createUserModel"
// @Success 200 {object} models.RespUser
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/user/update [put]
func (h *handlerV1) UpdateUser(c *gin.Context) {
	var (
		body        pbu.User
		jsonMarshal protojson.MarshalOptions
	)
	jsonMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", l.Error(err))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.UserService().UpdateUser(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to update user", l.Error(err))
		return
	}
	resp := models.RespUser{
		Id:        response.Id,
		FirstName: response.FirstName,
		LastName:  response.LastName,
		Bio:       response.Bio,
		Website:   response.Website,
		UserName:  response.Username,
		Email:     response.Email,
		Password:  response.Password,
	}

	c.JSON(http.StatusOK, resp)
}

// DeleteUser
// @Summary DeleteUser
// @Security ApiKeyAuth
// @Description Api for DeleteUser
// @Tags user
// @Accept json
// @Produce json
// @Param id query string true "ID"
// @Success 200 {object} models.User
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/user/delete [delete]
func (h *handlerV1) DeleteUser(c *gin.Context) {
	var jsonMarshal protojson.MarshalOptions
	jsonMarshal.UseProtoNames = true

	id := c.Query("id")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.UserService().DeleteUser(
		ctx, &pbu.UserId{
			UserId: id,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to delete user", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}
