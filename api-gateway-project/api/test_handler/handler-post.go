package test_handler

import (
	"log"
	"net/http"

	pbp "exam_task_4/api-gateway-project/genproto/post-service"
	u "exam_task_4/api-gateway-project/mock_data"
	"exam_task_4/api-gateway-project/storage"

	"github.com/gin-gonic/gin"
)

func CreatePost(ctx *gin.Context) {
	var newPost *storage.Post
	err := ctx.ShouldBindJSON(&newPost)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	mockClient := u.NewMockServiceClientP()

	res, err := mockClient.CreatePost(ctx, &pbp.Post{
		Id:       newPost.Id,
		UserId:   newPost.UserId,
		Title:    newPost.Title,
		Content:  newPost.Content,
		Category: newPost.Category,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create post"})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func GetPost(ctx *gin.Context) {
	postID := ctx.Query("id")
	mockClient := u.NewMockServiceClientP()
	response, err := mockClient.GetPost(ctx,
		&pbp.PostId{
			PostId: postID,
		})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get post"})
		log.Println("failed to get post")
		return
	}
	ctx.JSON(http.StatusOK, response)
}

func GetAllPosts(ctx *gin.Context) {
	mockClient := u.NewMockServiceClientP()
	response, err := mockClient.GetAllPost(ctx,
		&pbp.GetAllPostsRequest{
			Page:  2,
			Limit: 3,
		})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get posts"})
		log.Println("failed to get posts")
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func UpdatePost(ctx *gin.Context) {
	mockClient := u.NewMockServiceClientP()
	response, err := mockClient.UpdatePost(ctx,
		&pbp.Post{})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update post"})
		log.Println("failed to update post")
		return
	}
	ctx.JSON(http.StatusOK, response)
}

func DeletePost(ctx *gin.Context) {
	mockClient := u.NewMockServiceClientP()
	postID := ctx.Query("id")
	response, err := mockClient.DeletePost(
		ctx, &pbp.PostId{
			PostId: postID,
		})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, response)
}
