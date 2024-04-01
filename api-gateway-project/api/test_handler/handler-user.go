package test_handler

import (
	pbu "exam_task_4/api-gateway-project/genproto/user-service"
	u "exam_task_4/api-gateway-project/mock_data"
	"exam_task_4/api-gateway-project/storage"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(ctx *gin.Context) {
	var newUser *storage.User
	err := ctx.ShouldBindJSON(&newUser)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	mockClient := u.NewMockServiceClientU()

	res, err := mockClient.CreateUser(ctx, &pbu.User{
		Id:        newUser.Id,
		Username:  newUser.UserName,
		Email:     newUser.Email,
		Password:  newUser.Password,
		FirstName: newUser.FirstName,
		LastName:  newUser.LastName,
		Bio:       newUser.Bio,
		Website:   newUser.Website,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

// GetUser test handler
func GetUser(ctx *gin.Context) {
	userID := ctx.Query("id")
	mockClient := u.NewMockServiceClientU()
	response, err := mockClient.GetUser(ctx,
		&pbu.UserId{
			UserId: userID,
		})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get user"})
		log.Println("failed to get user")
		return
	}
	ctx.JSON(http.StatusOK, response)
}

// GetAll test handler
func GetAll(ctx *gin.Context) {
	mockClient := u.NewMockServiceClientU()
	response, err := mockClient.GetAllUser(ctx,
		&pbu.GetAllUsersRequest{
			Page:  2,
			Limit: 3,
		})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get user"})
		log.Println("failed to get user")
		return
	}

	ctx.JSON(http.StatusOK, response)
}

// Update test handler
func Update(ctx *gin.Context) {
	mockClient := u.NewMockServiceClientU()
	response, err := mockClient.UpdateUser(ctx,
		&pbu.User{})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get user"})
		log.Println("failed to get user")
		return
	}
	ctx.JSON(http.StatusOK, response)
}

// DeleteUser test handler
func DeleteUser(ctx *gin.Context) {
	mockClient := u.NewMockServiceClientU()
	userID := ctx.Query("id")
	response, err := mockClient.DeleteUser(
		ctx, &pbu.UserId{
			UserId: userID,
		})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, response)
}
