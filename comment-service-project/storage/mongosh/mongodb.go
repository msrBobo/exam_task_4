package mongo

import (
	"context"
	"errors"
	pbc "exam_task_4/comment-service-project/genproto/comment-service"
	m "exam_task_4/comment-service-project/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type commentRepo struct {
	collection *mongo.Collection
}

func NewCommentRepo(collection *mongo.Collection) *commentRepo {
	return &commentRepo{collection: collection}
}

func (c *commentRepo) CreateComment(comment *pbc.Comment) (*pbc.Comment, error) {
	comment.CreatedAt = time.Now().String()

	CreateComment := m.CreateComment{
		Id:        comment.Id,
		PostId:    comment.PostId,
		UserId:    comment.UserId,
		Content:   comment.Content,
		Likes:     comment.Likes,
		Dislikes:  comment.Dislikes,
		CreatedAt: comment.CreatedAt,
	}
	_, err := c.collection.InsertOne(context.Background(), CreateComment)
	if err != nil {
		return nil, err
	}

	return comment, nil
}

func (c *commentRepo) GetComment(Id *pbc.CommentId) (comment *pbc.Comment, err error) {
	filter := bson.M{"id": Id.Id}
	err = c.collection.FindOne(context.TODO(), filter).Decode(&comment)
	if err != nil {
		return nil, err
	}

	return comment, nil
}

func (c *commentRepo) GetAllComment(req *pbc.GetAllCommentsRequest) (*pbc.GetAllCommentsResponse, error) {

	ctx := context.TODO()

	offset := (req.Page - 1) * req.Limit

	findOptions := options.Find().
		SetSkip(int64(offset)).
		SetLimit(int64(req.Limit))

	filter := bson.M{"deleted_at": nil}

	cursor, err := c.collection.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var comments []*pbc.Comment
	for cursor.Next(context.Background()) {
		var comment pbc.Comment
		if err := cursor.Decode(&comment); err != nil {
			return nil, err
		}
		comments = append(comments, &comment)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return &pbc.GetAllCommentsResponse{Comments: comments}, nil
}

func (c *commentRepo) GetCommentsByPostId(post_id *pbc.PostId) (*pbc.GetAllCommentsResponse, error) {
	filter := bson.M{"postid": post_id.PostId, "deleted_at": nil}
	cursor, err := c.collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var response pbc.GetAllCommentsResponse
	for cursor.Next(context.Background()) {
		var comment pbc.Comment
		if err := cursor.Decode(&comment); err != nil {
			return nil, err
		}
		response.Comments = append(response.Comments, &comment)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *commentRepo) GetCommentsByUserId(post_id *pbc.PostId) (*pbc.GetAllCommentsResponse, error) {
	filter := bson.M{"userid": post_id.PostId, "deleted_at": nil}
	cursor, err := c.collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var response pbc.GetAllCommentsResponse
	for cursor.Next(context.Background()) {
		var comment pbc.Comment
		if err := cursor.Decode(&comment); err != nil {
			return nil, err
		}
		response.Comments = append(response.Comments, &comment)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *commentRepo) UpdateComment(comment *pbc.Comment) (*pbc.Comment, error) {
	filter := bson.M{"id": comment.Id}
	comment.UpdatedAt = time.Now().String()
	update := bson.M{
		"$set": bson.M{
			"content":    comment.Content,
			"likes":      comment.Likes,
			"dislikes":   comment.Dislikes,
			"updated_at": comment.UpdatedAt,
		},
	}
	_, err := c.collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return nil, err
	}

	var updatedComment pbc.Comment

	err = c.collection.FindOne(context.Background(), filter).Decode(&updatedComment)
	if err != nil {
		return nil, err
	}
	return &updatedComment, nil
}

func (c *commentRepo) DeleteComment(id *pbc.CommentId) (*pbc.DeleteResponse, error) {
	deletedAt := time.Now()
	filter := bson.M{"id": id.Id, "deleted_at": nil}
	update := bson.M{"$set": bson.M{"deleted_at": deletedAt}}
	result, err := c.collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return nil, err
	}
	if result.ModifiedCount == 0 {
		return nil, errors.New("delete failed: post not found or already deleted")
	}
	resp := &pbc.DeleteResponse{
		Message: "Deleted post",
	}
	return resp, nil
}
