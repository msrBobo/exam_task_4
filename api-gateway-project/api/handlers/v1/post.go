package v1

import (
	"context"
	"encoding/json"
	models "exam_task_4/api-gateway-project/api/handlers/models"
	pbp "exam_task_4/api-gateway-project/genproto/post-service"
	l "exam_task_4/api-gateway-project/pkg/logger"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"google.golang.org/protobuf/encoding/protojson"
)

// CreatePost
// @Summary CreatePost
// @Security ApiKeyAuth
// @Description Api for CreatePost
// @Tags post
// @Accept json
// @Produce json
// @Param Post body models.CreatePost true "createUserModel"
// @Success 200 {object} models.CreatePost
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/post/create [post]
func (h *handlerV1) CreatePost(c *gin.Context) {
	var (
		body        models.CreatePost
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

	// response, err := h.serviceManager.PostService().CreatePost(ctx, &pbp.Post{
	// 	Id:       body.Id,
	// 	UserId:   body.UserId,
	// 	Title:    body.Title,
	// 	Content:  body.Content,
	// 	Likes:    body.Likes,
	// 	Dislikes: body.Dislikes,
	// 	Views:    body.Views,
	// 	Category: body.Category,
	// })
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{
	// 		"error": err.Error(),
	// 	})
	// 	h.log.Error("failed to create post", l.Error(err))
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
	
	dataBytes, err := json.Marshal(&body)
	if err != nil {
		log.Println(err)
	}

	h.Writer.ProducerMessage("createPost", dataBytes)
	resp := models.CreatePost{
		Id:       body.Id,
		UserId:   body.UserId,
		Title:    body.Title,
		Content:  body.Content,
		Likes:    body.Likes,
		Dislikes: body.Dislikes,
		Views:    body.Views,
		Category: body.Category,
	}

	c.JSON(http.StatusCreated, resp)
}

// GetPost
// @Summary GetPost
// @Security ApiKeyAuth
// @Description Api for GetPost
// @Tags post
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Success 200 {object} models.CreatePost
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/post/{id} [get]
func (h *handlerV1) GetPost(c *gin.Context) {
	var jsonMarshal protojson.MarshalOptions
	jsonMarshal.UseProtoNames = true

	id := c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.PostService().GetPost(
		ctx, &pbp.PostId{
			PostId: id,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get post", l.Error(err))
		return
	}

	resp := models.CreatePost{
		Id:       response.Id,
		UserId:   response.UserId,
		Title:    response.Title,
		Content:  response.Content,
		Likes:    response.Likes,
		Dislikes: response.Dislikes,
		Views:    response.Views,
		Category: response.Category,
	}

	c.JSON(http.StatusOK, resp)
}

// GetAllPosts
// @Summary GetAllPosts
// @Security ApiKeyAuth
// @Description Api for GetAllPosts
// @Tags post
// @Accept json
// @Produce json
// @Param page query string true "User PAGES"
// @Param limit query string true "User LIMIT"
// @Success 200 {object} models.CreatePost
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/posts [get]
func (h *handlerV1) GetAllPosts(c *gin.Context) {
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

	response, err := h.serviceManager.PostService().GetAllPost(
		ctx, &pbp.GetAllPostsRequest{
			Limit: int64(LimitNum),
			Page:  int64(PageNum),
		})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to list posts", l.Error(err))
		return
	}
	var respUsers []models.CreatePost
	for _, post := range response.Posts {
		resp := models.CreatePost{
			Id:       post.Id,
			UserId:   post.UserId,
			Title:    post.Title,
			Content:  post.Content,
			Likes:    post.Likes,
			Dislikes: post.Dislikes,
			Views:    post.Views,
			Category: post.Category,
		}
		respUsers = append(respUsers, resp)
	}

	c.JSON(http.StatusOK, respUsers)
}

// GetPostsByUserId
// @Summary GetPostsByUserId
// @Security ApiKeyAuth
// @Description Api for GetPostsByUserId
// @Tags post
// @Accept json
// @Produce json
// @Param id path string true "Id"
// @Success 200 {object} models.CreatePost
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/postsbyuserid/{id} [get]
func (h *handlerV1) GetPostsByUserId(c *gin.Context) {
	var jsonMarshal protojson.MarshalOptions
	jsonMarshal.UseProtoNames = true

	id := c.Param("id")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.PostService().GetPostsByUserId(
		ctx, &pbp.UserId{
			UserId: id,
		})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to list posts", l.Error(err))
		return
	}
	var respUsers []models.CreatePost
	for _, post := range response.Posts {
		resp := models.CreatePost{
			Id:       post.Id,
			UserId:   post.UserId,
			Title:    post.Title,
			Content:  post.Content,
			Likes:    post.Likes,
			Dislikes: post.Dislikes,
			Views:    post.Views,
			Category: post.Category,
		}
		respUsers = append(respUsers, resp)
	}

	c.JSON(http.StatusOK, respUsers)
}

// UpdatePost
// @Summary UpdatePost
// @Security ApiKeyAuth
// @Description Api for UpdatePost
// @Tags post
// @Accept json
// @Produce json
// @Param Post body models.Post true "createUserModel"
// @Success 200 {object} models.CreatePost
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/post/update [put]
func (h *handlerV1) UpdatePost(c *gin.Context) {
	var (
		body        pbp.Post
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

	response, err := h.serviceManager.PostService().UpdatePost(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to update post", l.Error(err))
		return
	}
	resp := models.CreatePost{
		Id:       response.Id,
		UserId:   response.UserId,
		Title:    response.Title,
		Content:  response.Content,
		Likes:    response.Likes,
		Dislikes: response.Dislikes,
		Views:    response.Views,
		Category: response.Category,
	}

	c.JSON(http.StatusOK, resp)
}

// DeletePost
// @Summary DeletePost
// @Security ApiKeyAuth
// @Description Api for DeletePost
// @Tags post
// @Accept json
// @Produce json
// @Param id query string true "ID"
// @Success 200 {object} models.Post
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/post/delete [delete]
func (h *handlerV1) DeletePost(c *gin.Context) {
	var jsonMarshal protojson.MarshalOptions
	jsonMarshal.UseProtoNames = true

	id := c.Query("id")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.PostService().DeletePost(
		ctx, &pbp.PostId{
			PostId: id,
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
