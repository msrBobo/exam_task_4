package mongosh

import (
	"context"
	"errors"
	pbp "exam_task_4/post-service-project/genproto/post-service"
	m "exam_task_4/post-service-project/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type postRepo struct {
	collection *mongo.Collection
}

func NewPostRepo(collection *mongo.Collection) *postRepo {
	return &postRepo{collection: collection}
}

func (p *postRepo) CreatePost(post *pbp.Post) (*pbp.Post, error) {
	post.CreatedAt = time.Now().String()
	CreatePosts := m.CreatePost{
		Id:        post.Id,
		UserID:    post.UserId,
		Title:     post.Title,
		Content:   post.Content,
		Likes:     post.Likes,
		Dislikes:  post.Dislikes,
		Views:     post.Views,
		Category:  post.Category,
		CreatedAt: post.CreatedAt,
	}
	_, err := p.collection.InsertOne(context.Background(), CreatePosts)
	if err != nil {
		return nil, err
	}

	return post, nil
}

func (p *postRepo) GetPost(Id *pbp.PostId) (post *pbp.Post, err error) {
	filter := bson.M{"id": Id.PostId}
	err = p.collection.FindOne(context.TODO(), filter).Decode(&post)
	if err != nil {
		return nil, err
	}

	return post, nil
}

func (p *postRepo) GetAllPost(req *pbp.GetAllPostsRequest) (*pbp.GetAllPostsResponse, error) {
	ctx := context.TODO()

	offset := (req.Page - 1) * req.Limit

	findOptions := options.Find().
		SetSkip(int64(offset)).
		SetLimit(int64(req.Limit))

	filter := bson.M{"deleted_at": nil}

	cursor, err := p.collection.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var posts []*pbp.Post
	for cursor.Next(ctx) {
		var post pbp.Post
		if err := cursor.Decode(&post); err != nil {
			return nil, err
		}
		posts = append(posts, &post)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return &pbp.GetAllPostsResponse{Posts: posts}, nil
}

func (p *postRepo) GetPostsByUserId(id *pbp.UserId) (*pbp.GetAllPostsResponse, error) {
	filter := bson.M{"userid": id.UserId, "deleted_at": nil}
	cursor, err := p.collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var response pbp.GetAllPostsResponse
	for cursor.Next(context.Background()) {
		var post pbp.Post
		err := cursor.Decode(&post)
		if err != nil {
			return nil, err
		}
		response.Posts = append(response.Posts, &post)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return &response, nil
}

func (p *postRepo) UpdatePost(post *pbp.Post) (*pbp.Post, error) {
	filter := bson.M{"id": post.Id}
	post.UpdatedAt = time.Now().String()
	update := bson.M{
		"$set": bson.M{
			"content":    post.Content,
			"title":      post.Title,
			"updated_at": time.Now(),
			"likes":      post.Likes,
			"dislikes":   post.Dislikes,
			"views":      post.Views,
			"category":   post.Category,
			"updatedat":  post.UpdatedAt,
		},
	}
	_, err := p.collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return nil, err
	}

	var updatedPost pbp.Post
	err = p.collection.FindOne(context.Background(), filter).Decode(&updatedPost)
	if err != nil {
		return nil, err
	}

	return &updatedPost, nil
}

func (p *postRepo) DeletePost(postId *pbp.PostId) (*pbp.DeleteResponse, error) {
	deletedAt := time.Now()
	filter := bson.M{"id": postId.PostId, "deleted_at": nil}
	update := bson.M{"$set": bson.M{"deleted_at": deletedAt}}
	result, err := p.collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return nil, err
	}
	if result.ModifiedCount == 0 {
		return nil, errors.New("delete failed: post not found or already deleted")
	}
	resp := &pbp.DeleteResponse{
		Message: "Deleted post",
	}
	return resp, nil
}
