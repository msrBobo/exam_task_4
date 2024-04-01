package v1

import (
	"context"
	"encoding/json"
	models "exam_task_4/api-gateway-project/api/handlers/models"
	pbc "exam_task_4/api-gateway-project/genproto/comment-service"
	l "exam_task_4/api-gateway-project/pkg/logger"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"google.golang.org/protobuf/encoding/protojson"
)

// CreateComment
// @Summary CreateComment
// @Security ApiKeyAuth
// @Description Api for CreateComment
// @Tags comment
// @Accept json
// @Produce json
// @Param Comment body models.CreateComment true "createCommentModel"
// @Success 200 {object} models.CreateComment
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/comment/create [post]
func (h *handlerV1) CreateComment(c *gin.Context) {

	var (
		body        models.CreateComment
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

	// ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	// defer cancel()

	// response, err := h.serviceManager.CommentService().CreateComment(ctx, &pbc.Comment{
	// 	PostId:   body.PostId,
	// 	UserId:   body.UserId,
	// 	Content:  body.Content,
	// 	Likes:    body.Likes,
	// 	Dislikes: body.Dislikes,
	// })
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{
	// 		"error": err.Error(),
	// 	})
	// 	h.log.Error("failed to create user", l.Error(err))
	// 	return
	// }
	if body.Id  == "" || body.Id == "string" {
		id := uuid.New().String()
		body.Id = id
	}

	if body.UserId  == "" || body.UserId == "string" {
		id := uuid.New().String()
		body.UserId = id
	}
	
	if body.PostId  == "" || body.PostId == "string" {
		id := uuid.New().String()
		body.PostId = id
	}
	
	dataBytes, err := json.Marshal(&body)
	if err != nil {
		log.Println(err)
	}

	h.Writer.ProducerMessage("createComment", dataBytes)

	resp := models.CreateComment{
		Id:       body.Id,
		PostId:   body.PostId,
		UserId:   body.UserId,
		Content:  body.Content,
		Likes:    body.Likes,
		Dislikes: body.Dislikes,
	}

	c.JSON(http.StatusCreated, resp)
}

// GetComment
// @Summary GetComment
// @Security ApiKeyAuth
// @Description Api for GetComment
// @Tags comment
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Success 200 {object} models.CreateComment
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/comment/{id} [get]
func (h *handlerV1) GetComment(c *gin.Context) {
	var jsonMarshal protojson.MarshalOptions
	jsonMarshal.UseProtoNames = true

	id := c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.CommentService().GetComment(
		ctx, &pbc.CommentId{
			Id: id,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get user", l.Error(err))
		return
	}

	resp := models.CreateComment{
		Id:       response.Id,
		PostId:   response.PostId,
		UserId:   response.UserId,
		Content:  response.Content,
		Likes:    response.Likes,
		Dislikes: response.Dislikes,
	}

	c.JSON(http.StatusOK, resp)
}

// GetAllComments
// @Summary GetAllComments
// @Security ApiKeyAuth
// @Description Api for GetAllComments
// @Tags comment
// @Accept json
// @Produce json
// @Param page query string true "User PAGES"
// @Param limit query string true "User LIMIT"
// @Success 200 {object} models.CreateComment
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/comments [get]
func (h *handlerV1) GetAllComments(c *gin.Context) {
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
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.CommentService().GetAllComment(
		ctx, &pbc.GetAllCommentsRequest{
			Limit: int64(LimitNum),
			Page:  int64(PageNum),
		})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to list users", l.Error(err))
		return
	}

	var respUsers []models.CreateComment
	for _, comment := range response.Comments {
		resp := models.CreateComment{
			Id:       comment.Id,
			PostId:   comment.PostId,
			UserId:   comment.UserId,
			Content:  comment.Content,
			Likes:    comment.Likes,
			Dislikes: comment.Dislikes,
		}
		respUsers = append(respUsers, resp)
	}

	c.JSON(http.StatusOK, respUsers)
}

// GetCommentsByPostId
// @Summary GetCommentsByPostId
// @Security ApiKeyAuth
// @Description Api for GetCommentsByPostId
// @Tags comment
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Success 200 {object} models.CreateComment
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/commentsby/{id} [get]
func (h *handlerV1) GetCommentsByPostId(c *gin.Context) {
	var jsonMarshal protojson.MarshalOptions
	jsonMarshal.UseProtoNames = true

	id := c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.CommentService().GetCommentsByPostId(
		ctx, &pbc.PostId{
			PostId: id,
		})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to list comments", l.Error(err))
		return
	}

	var respComments []models.CreateComment
	for _, comment := range response.Comments {
		resp := models.CreateComment{
			Id:       comment.Id,
			PostId:   comment.PostId,
			UserId:   comment.UserId,
			Content:  comment.Content,
			Likes:    comment.Likes,
			Dislikes: comment.Dislikes,
		}
		respComments = append(respComments, resp)
	}

	c.JSON(http.StatusOK, respComments)
}

// GetCommentsByUserId
// @Summary GetCommentsByUserId
// @Security ApiKeyAuth
// @Description Api for GetCommentsByUserId
// @Tags comment
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Success 200 {object} models.CreateComment
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/commentsbyuser/{id} [get]
func (h *handlerV1) GetCommentsByUserId(c *gin.Context) {
	var jsonMarshal protojson.MarshalOptions
	jsonMarshal.UseProtoNames = true

	id := c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.CommentService().GetCommentsByUserId(
		ctx, &pbc.PostId{
			PostId: id,
		})

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to list comments", l.Error(err))
		return
	}

	var respComments []models.CreateComment
	for _, comment := range response.Comments {
		resp := models.CreateComment{
			Id:       comment.Id,
			PostId:   comment.PostId,
			UserId:   comment.UserId,
			Content:  comment.Content,
			Likes:    comment.Likes,
			Dislikes: comment.Dislikes,
		}
		respComments = append(respComments, resp)
	}

	c.JSON(http.StatusOK, respComments)
}

// UpdateComment
// @Summary UpdateComment
// @Security ApiKeyAuth
// @Description Api for UpdateComment
// @Tags comment
// @Accept json
// @Produce json
// @Param Comment body models.CreateComment true "createCommentModel"
// @Success 200 {object} models.CreateComment
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/comment/update [put]
func (h *handlerV1) UpdateComment(c *gin.Context) {
	var (
		body        pbc.Comment
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

	response, err := h.serviceManager.CommentService().UpdateComment(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to update comment", l.Error(err))
		return
	}
	resp := models.CreateComment{
		Id:       response.Id,
		PostId:   response.PostId,
		UserId:   response.UserId,
		Content:  response.Content,
		Likes:    response.Likes,
		Dislikes: response.Dislikes,
	}

	c.JSON(http.StatusOK, resp)
}

// DeleteComment
// @Summary DeleteComment
// @Security ApiKeyAuth
// @Description Api for DeleteComment
// @Tags comment
// @Accept json
// @Produce json
// @Param id query string true "ID"
// @Success 200 {object} models.CreateComment
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/comment/delete [delete]
func (h *handlerV1) DeleteComment(c *gin.Context) {
	var jsonMarshal protojson.MarshalOptions
	jsonMarshal.UseProtoNames = true

	id := c.Query("id")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.CommentService().DeleteComment(
		ctx, &pbc.CommentId{
			Id: id,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to delete comment", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}
