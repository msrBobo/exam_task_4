package v1

import (
	"context"
	models "exam_task_4/api-gateway-project/api/handlers/models"
	pbc "exam_task_4/api-gateway-project/genproto/comment-service"
	pbp "exam_task_4/api-gateway-project/genproto/post-service"
	pbu "exam_task_4/api-gateway-project/genproto/user-service"
	l "exam_task_4/api-gateway-project/pkg/logger"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// GetPostsByUserIdWithComments
// @Summary GetPostsByUserIdWithComments
// @Security ApiKeyAuth
// @Description Api for GetPostsByUserIdWithComments
// @Tags GetAllPostsWithCommentsAndOwners
// @Accept json
// @Produce json
// @Param id path string true "Id"
// @Success 200 {object} models.CreatePost
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/postsbyuseridwithcomments/{id} [get]
// GetPostsAndCommentsByUserId retrieves posts and comments by user ID
// GetPostsAndCommentsByUserId retrieves posts and comments by user ID
func (h *handlerV1) GetPostsAndCommentsByUserId(c *gin.Context) {
	id := c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	postsResponse, err := h.serviceManager.PostService().GetPostsByUserId(
		ctx, &pbp.UserId{
			UserId: id,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		h.log.Error("failed to list posts", l.Error(err))
		return
	}
	var userId string
	var respPosts []models.RespPost
	for _, post := range postsResponse.Posts {
		resp := models.RespPost{
			Title:    post.Title,
			Content:  post.Content,
			Likes:    post.Likes,
			Dislikes: post.Dislikes,
			Views:    post.Views,
			Category: post.Category,
		}
		respPosts = append(respPosts, resp)
		userId = post.UserId
	}
	fmt.Println(userId)
	commentsResponse, err := h.serviceManager.CommentService().GetCommentsByUserId(
		ctx, &pbc.PostId{
			PostId: userId,
		})
	if err != nil {
		return
	}

	var respComments []models.RespComment
	for _, comment := range commentsResponse.Comments {
		resp := models.RespComment{
			Content:  comment.Content,
			Likes:    comment.Likes,
			Dislikes: comment.Dislikes,
		}
		respComments = append(respComments, resp)
	}

	c.JSON(http.StatusOK, gin.H{
		"posts":    respPosts,
		"comments": respComments,
	})
}

// GetAllPostsWithCommentsAndOwners retrieves all posts with comments and owners
// @Summary Retrieves all posts with comments and owners
// @Security ApiKeyAuth
// @Description Retrieves all posts with comments and owners
// @Tags GetAllPostsWithCommentsAndOwners
// @Accept json
// @Produce json
// @Param page query string true "User PAGES"
// @Param limit query string true "User LIMIT"
// @Success 200 {object} models.AllPostWithCommentsAndOwnersResponse
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/allpostswithcommentsandowners [get]
func (h *handlerV1) GetAllPostsWithCommentsAndOwners(c *gin.Context) {

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

	postsResponse, err := h.serviceManager.PostService().GetAllPost(ctx, &pbp.GetAllPostsRequest{Page: int64(PageNum), Limit: int64(LimitNum)})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		h.log.Error("failed to list posts", l.Error(err))
		return
	}

	commentsResponse, err := h.serviceManager.CommentService().GetAllComment(ctx, &pbc.GetAllCommentsRequest{Page: int64(PageNum), Limit: int64(LimitNum)})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		h.log.Error("failed to list comments", l.Error(err))
		return
	}

	usersResponse, err := h.serviceManager.UserService().GetAllUsers(ctx, &pbu.GetAllUsersRequest{Page: int64(PageNum), Limit: int64(LimitNum)})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		h.log.Error("failed to list users", l.Error(err))
		return
	}

	var respPosts []models.AllPostWithCommentsAndOwnersResponse
	for _, post := range postsResponse.Posts {
		respPost := models.AllPostWithCommentsAndOwnersResponse{
			Posts: []*models.Post{{
				Id:       post.Id,
				UserId:   post.UserId,
				Title:    post.Title,
				Content:  post.Content,
				Likes:    post.Likes,
				Dislikes: post.Dislikes,
				Views:    post.Views,
				Category: post.Category,
			}},
		}

		for _, comment := range commentsResponse.Comments {
			if comment.PostId == post.Id {
				respPost.Comments = append(respPost.Comments, &models.Comment{
					Id:       comment.Id,
					PostId:   comment.PostId,
					UserId:   comment.UserId,
					Content:  comment.Content,
					Likes:    comment.Likes,
					Dislikes: comment.Dislikes,
				})
			}
		}

		for _, user := range usersResponse.Users {
			if user.Id == post.UserId {
				respPost.Owner = &models.User{
					Id:        user.Id,
					UserName:  user.Username,
					Email:     user.Email,
					Password:  user.Password,
					FirstName: user.FirstName,
					LastName:  user.LastName,
					Bio:       user.Bio,
					Website:   user.Website,
				}
				break
			}
		}

		respPosts = append(respPosts, respPost)
	}

	c.JSON(http.StatusOK, respPosts)
}
