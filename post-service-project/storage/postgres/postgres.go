package postgres

import (
	pbp "exam_task_4/post-service-project/genproto/post-service"
	"time"

	"github.com/jmoiron/sqlx"
)

type postRepoPostgres struct {
	db *sqlx.DB
}

func NewPostRepoPostgres(db *sqlx.DB) *postRepoPostgres {
	return &postRepoPostgres{db: db}
}

func (p *postRepoPostgres) CreatePost(post *pbp.Post) (*pbp.Post, error) {
	CreatedAt := time.Now()
	err := p.db.QueryRow(`
		INSERT INTO posts
			(
				id,
				user_id,							
				content,
				title,
				created_at,
				likes,
				dislikes,
				views,
				category
			)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
			RETURNING
				id,
				user_id,
				content,
				title,
				likes,
				dislikes,
				views,
				category`,
		post.Id,
		post.UserId,
		post.Content,
		post.Title,
		CreatedAt,
		post.Likes,
		post.Dislikes,
		post.Views,
		post.Category,
	).Scan(
		&post.Id,
		&post.UserId,
		&post.Content,
		&post.Title,
		&post.Likes,
		&post.Dislikes,
		&post.Views,
		&post.Category,
	)
	if err != nil {
		return nil, err
	}

	return post, nil
}

func (p *postRepoPostgres) GetPost(postId *pbp.PostId) (post *pbp.Post, err error) {
	post = &pbp.Post{}

	err = p.db.QueryRow(`
        SELECT
            id,
            user_id,
            content,
            title,
            likes,
            dislikes,
            views,
            category
        FROM posts
        WHERE id = $1 AND deleted_at IS NULL`,
		postId.PostId,
	).Scan(
		&post.Id,
		&post.UserId,
		&post.Content,
		&post.Title,
		&post.Likes,
		&post.Dislikes,
		&post.Views,
		&post.Category,
	)

	if err != nil {
		return nil, err
	}

	return post, nil
}

func (p *postRepoPostgres) GetAllPost(req *pbp.GetAllPostsRequest) (*pbp.GetAllPostsResponse, error) {
	offset := (req.Page - 1) * req.Limit
	rows, err := p.db.Query(`
		SELECT
			id,
			user_id,
			content,
			title,
			likes,
			dislikes,
			views,
			category
		FROM posts
		WHERE deleted_at IS NULL
		LIMIT $1 OFFSET $2
	`, req.Limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var response pbp.GetAllPostsResponse
	for rows.Next() {
		var post pbp.Post
		err := rows.Scan(
			&post.Id,
			&post.UserId,
			&post.Content,
			&post.Title,
			&post.Likes,
			&post.Dislikes,
			&post.Views,
			&post.Category,
		)
		if err != nil {
			return nil, err
		}
		response.Posts = append(response.Posts, &post)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return &response, nil
}

func (p *postRepoPostgres) GetPostsByUserId(id *pbp.UserId) (*pbp.GetAllPostsResponse, error) {
	rows, err := p.db.Query(`
			SELECT
				id,
				user_id,
				content,
				title,
				likes,
				dislikes,
				views,
				category
			FROM posts
			WHERE user_id = $1 AND deleted_at IS NULL`, id.UserId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var response pbp.GetAllPostsResponse
	for rows.Next() {
		var post pbp.Post
		err := rows.Scan(
			&post.Id,
			&post.UserId,
			&post.Title,
			&post.Content,
			&post.Likes,
			&post.Dislikes,
			&post.Views,
			&post.Category,
		)
		if err != nil {
			return nil, err
		}
		response.Posts = append(response.Posts, &post)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return &response, nil
}

func (p *postRepoPostgres) UpdatePost(post *pbp.Post) (*pbp.Post, error) {

	UpdatedAt := time.Now()
	_, err := p.db.Exec(`
		UPDATE posts
			SET
				content = $1,
				title = $2,
				update_at = $3,
				likes = $4,
				dislikes = $5,
				views = $6,
				category = $7
			WHERE id = $8 AND deleted_at IS NULL`,
		post.Content,
		post.Title,
		UpdatedAt,
		post.Likes,
		post.Dislikes,
		post.Views,
		post.Category,
		post.Id)

	if err != nil {
		return nil, err
	}
	id := &pbp.PostId{
		PostId: post.Id,
	}
	resp, err := p.GetPost(id)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *postRepoPostgres) DeletePost(postId *pbp.PostId) (*pbp.DeleteResponse, error) {
	resp := &pbp.DeleteResponse{}

	deletedAt := time.Now()

	result, err := p.db.Exec(`
		UPDATE posts
		SET deleted_at = $1
		WHERE id = $2 AND deleted_at IS NULL`, deletedAt, postId.PostId)
	if err != nil {
		return nil, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rowsAffected == 0 {
		resp.Message = "Post not found or already deleted"
		return resp, nil
	}
	resp.Message = "Deleted post"

	return resp, nil
}
