package tests

import (
	"encoding/json"
	handler "exam_task_4/api-gateway-project/api/test_handler"
	"exam_task_4/api-gateway-project/storage"

	"net/http"
	"testing"

	"github.com/gin-gonic/gin"

	"github.com/k0kubun/pp"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestApiUser(t *testing.T) {
	require.NoError(t, SetupMinimumInstance(""))
	buf, err := OpenFile("user.json")
	require.NoError(t, err)

	req := NewRequest(http.MethodPost, "/user/create", buf)

	router := gin.New()
	res, err := Server(handler.CreateUser, "/user/create", req, router)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.Code)
	var user *storage.User
	require.NoError(t, json.Unmarshal(res.Body.Bytes(), &user))
	require.Equal(t, user.FirstName, "testfirstname")
	require.Equal(t, user.LastName, "testlastname")
	require.Equal(t, user.Email, "testemail")
	require.Equal(t, user.Password, "testpassword")
	require.NotNil(t, user.Id)

	//GetUser
	getReq := NewRequest(http.MethodGet, "/user/get", nil)
	args := getReq.URL.Query()
	args.Add("id", user.Id)
	getReq.URL.RawQuery = args.Encode()
	getRes, err := Server(handler.GetUser, "/user/get", getReq, router)
	assert.NoError(t, err)
	require.Equal(t, user.FirstName, "testfirstname")
	require.Equal(t, user.LastName, "testlastname")
	require.Equal(t, user.Email, "testemail")
	require.Equal(t, user.Password, "testpassword")
	require.NotNil(t, user.Id)
	assert.Equal(t, http.StatusOK, getRes.Code)
	pp.Print(user)

	//GetAll
	getsReq := NewRequest(http.MethodGet, "/user/getall", nil)
	getsRes, err := Server(handler.GetAll, "/user/getall", getsReq, router)
	assert.NoError(t, err)
	require.Equal(t, user.FirstName, "testfirstname")
	require.Equal(t, user.LastName, "testlastname")
	require.Equal(t, user.Email, "testemail")
	require.Equal(t, user.Password, "testpassword")
	require.NotNil(t, user.Id)
	assert.Equal(t, http.StatusOK, getsRes.Code)
	var getsUser *storage.User
	require.NoError(t, json.Unmarshal(getRes.Body.Bytes(), &getsUser))

	//UpdUser
	getReqU := NewRequest(http.MethodPut, "/user/update", buf)
	getResU, err := Server(handler.Update, "/user/update", getReqU, router)
	assert.NoError(t, err)
	require.Equal(t, user.FirstName, "testfirstname")
	require.Equal(t, user.LastName, "testlastname")
	require.Equal(t, user.Email, "testemail")
	require.Equal(t, user.Password, "testpassword")
	require.NotNil(t, user.Id)
	assert.Equal(t, http.StatusOK, getResU.Code)
	require.NoError(t, json.Unmarshal(getRes.Body.Bytes(), &user))

	//DeleteUser
	deleteReq := NewRequest(http.MethodDelete, "/user/delete", buf)
	args = deleteReq.URL.Query()
	args.Add("id", user.Id)
	deleteReq.URL.RawQuery = args.Encode()

	deleteRes, err := Server(handler.DeleteUser, "/user/delete", deleteReq, router)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, deleteRes.Code)
	var deleteUser *storage.User
	require.NoError(t, json.Unmarshal(res.Body.Bytes(), &deleteUser))
	require.Equal(t, user.Id, deleteUser.Id)
}

func TestApiPost(t *testing.T) {
	require.NoError(t, SetupMinimumInstance(""))
	buf, err := OpenFile("post.json")
	require.NoError(t, err)

	// Create Post
	reqCreatePost := NewRequest(http.MethodPost, "/post/create", buf)
	router := gin.New()
	resCreatePost, err := Server(handler.CreatePost, "/post/create", reqCreatePost, router)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resCreatePost.Code)

	// Get Post
	var createdPost *storage.Post
	require.NoError(t, json.Unmarshal(resCreatePost.Body.Bytes(), &createdPost))
	getReq := NewRequest(http.MethodGet, "/post/get", nil)
	args := getReq.URL.Query()
	args.Add("id", createdPost.Id)
	getReq.URL.RawQuery = args.Encode()
	resGetPost, err := Server(handler.GetPost, "/post/get", getReq, router)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resGetPost.Code)
	var getPost *storage.Post
	require.NoError(t, json.Unmarshal(resGetPost.Body.Bytes(), &getPost))
	assert.Equal(t, createdPost.Id, getPost.Id)
	pp.Print(createdPost)
	// // Get All Posts
	reqGetAllPosts := NewRequest(http.MethodGet, "/post/getall", nil)
	resGetAllPosts, err := Server(handler.GetAllPosts, "/post/getall", reqGetAllPosts, router)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resGetAllPosts.Code)

	// // Update Post
	reqUpdatePost := NewRequest(http.MethodPut, "/post/update", buf)
	resUpdatePost, err := Server(handler.UpdatePost, "/post/update", reqUpdatePost, router)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resUpdatePost.Code)

	// // Delete Post
	reqDeletePost := NewRequest(http.MethodDelete, "/post/delete", buf)
	args = reqDeletePost.URL.Query()
	args.Add("id", createdPost.Id)
	reqDeletePost.URL.RawQuery = args.Encode()
	resDeletePost, err := Server(handler.DeletePost, "/post/delete", reqDeletePost, router)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resDeletePost.Code)
}

// TestApiComment tests the comment API endpoints.
func TestApiComment(t *testing.T) {
	require.NoError(t, SetupMinimumInstance(""))

	buf, err := OpenFile("comment.json")
	require.NoError(t, err)

	// Create Comment
	reqCreateComment := NewRequest(http.MethodPost, "/comment/create", buf)
	router := gin.New()
	resCreateComment, err := Server(handler.CreateComment, "/comment/create", reqCreateComment, router)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resCreateComment.Code)

	var createdComment *storage.Comment
	require.NoError(t, json.Unmarshal(resCreateComment.Body.Bytes(), &createdComment))

	// Get Comment
	getReq := NewRequest(http.MethodGet, "/comment/get", nil)
	args := getReq.URL.Query()
	args.Add("id", createdComment.Id)
	getReq.URL.RawQuery = args.Encode()
	resGetComment, err := Server(handler.GetComment, "/comment/get", getReq, router)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resGetComment.Code)
	pp.Println(createdComment)
	// Update Comment (if applicable)
	updateReq := NewRequest(http.MethodPut, "/comment/update", buf)
	updateRes, err := Server(handler.UpdateComment, "/comment/update", updateReq, router)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, updateRes.Code)

	// Delete Comment (if applicable)
	deleteReq := NewRequest(http.MethodDelete, "/comment/delete", buf)
	deleteRes, err := Server(handler.DeleteComment, "/comment/delete", deleteReq, router)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, deleteRes.Code)
}
