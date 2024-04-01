package test_handler

import (
	"log"
	"net/http"

	pbc "exam_task_4/api-gateway-project/genproto/comment-service"
	u "exam_task_4/api-gateway-project/mock_data"

	"github.com/gin-gonic/gin"
)

func CreateComment(ctx *gin.Context) {
	var newComment *pbc.Comment
	err := ctx.ShouldBindJSON(&newComment)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	mockClient := u.NewMockServiceClientC()

	res, err := mockClient.CreateComment(ctx, &pbc.Comment{
		Id:       newComment.Id,
		PostId:   newComment.PostId,
		UserId:   newComment.UserId,
		Content:  newComment.Content,
		Likes:    newComment.Likes,
		Dislikes: newComment.Dislikes,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create comment"})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func GetComment(ctx *gin.Context) {
	commentID := ctx.Query("id")
	mockClient := u.NewMockServiceClientC()
	response, err := mockClient.GetComment(ctx,
		&pbc.CommentId{
			Id: commentID,
		})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get comment"})
		log.Println("failed to get comment")
		return
	}
	ctx.JSON(http.StatusOK, response)
}

func GetAllComments(ctx *gin.Context) {
	mockClient := u.NewMockServiceClientC()
	response, err := mockClient.GetAllComment(ctx,
		&pbc.GetAllCommentsRequest{})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get comments"})
		log.Println("failed to get comments")
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func UpdateComment(ctx *gin.Context) {
	mockClient := u.NewMockServiceClientC()
	response, err := mockClient.UpdateComment(ctx,
		&pbc.Comment{})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update comment"})
		log.Println("failed to update comment")
		return
	}
	ctx.JSON(http.StatusOK, response)
}

func DeleteComment(ctx *gin.Context) {
	mockClient := u.NewMockServiceClientC()
	commentID := ctx.Query("id")
	response, err := mockClient.DeleteComment(
		ctx, &pbc.CommentId{
			Id: commentID,
		})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, response)
}
