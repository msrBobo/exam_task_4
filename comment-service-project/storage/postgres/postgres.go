package postgres

import (
	pbc "exam_task_4/comment-service-project/genproto/comment-service"
	"time"

	"github.com/jmoiron/sqlx"
)

type commentRepoPostgres struct {
	db *sqlx.DB
}

func NewCommentRepoPostgres(db *sqlx.DB) *commentRepoPostgres {
	return &commentRepoPostgres{db: db}
}

func (c *commentRepoPostgres) CreateComment(comment *pbc.Comment) (*pbc.Comment, error) {

	CreatedAt := time.Now()
	err := c.db.QueryRow(`
		INSERT INTO comments
			(
				id,
				post_id,
				user_id,
				content,
				likes,
				dislikes,
				created_at
			)
			VALUES ($1, $2, $3, $4, $5, $6, $7)
			RETURNING
				id,
				post_id,
				user_id,
				content, 
				likes,
				dislikes`,
		comment.Id,
		comment.PostId,
		comment.UserId,
		comment.Content,
		comment.Likes,
		comment.Dislikes,
		CreatedAt,
	).Scan(
		&comment.Id,
		&comment.PostId,
		&comment.UserId,
		&comment.Content,
		&comment.Likes,
		&comment.Dislikes,
	)
	if err != nil {
		return nil, err
	}

	return comment, nil
}

func (c *commentRepoPostgres) GetComment(id *pbc.CommentId) (*pbc.Comment, error) {
	comment := &pbc.Comment{}
	err := c.db.QueryRow(`
		SELECT
			id,
			post_id,
			user_id,
			content,
			likes,
			dislikes
		FROM comments
		WHERE id = $1 AND deleted_at IS NULL`,
		id.Id,
	).Scan(
		&comment.Id,
		&comment.PostId,
		&comment.UserId,
		&comment.Content,
		&comment.Likes,
		&comment.Dislikes,
	)
	if err != nil {
		return nil, err
	}

	return comment, nil
}

func (c *commentRepoPostgres) GetAllComment(req *pbc.GetAllCommentsRequest) (*pbc.GetAllCommentsResponse, error) {
	offset := (req.Page - 1) * req.Limit
	rows, err := c.db.Query(`
		SELECT
			id,
			post_id,
			user_id,
			content,
			likes,
			dislikes
		FROM comments
		WHERE deleted_at IS NULL
		LIMIT $1 OFFSET $2`, req.Limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var response pbc.GetAllCommentsResponse
	for rows.Next() {
		var comment pbc.Comment
		err := rows.Scan(
			&comment.Id,
			&comment.PostId,
			&comment.UserId,
			&comment.Content,
			&comment.Likes,
			&comment.Dislikes,
		)
		if err != nil {
			return nil, err
		}
		response.Comments = append(response.Comments, &comment)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *commentRepoPostgres) GetCommentsByPostId(post_id *pbc.PostId) (*pbc.GetAllCommentsResponse, error) {
	rows, err := c.db.Query(`
		SELECT
			id,
			post_id,
			user_id,
			content,
			likes,
			dislikes
		FROM comments
		WHERE post_id = $1 AND deleted_at IS NULL`, post_id.PostId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var response pbc.GetAllCommentsResponse
	for rows.Next() {
		var comment pbc.Comment
		err := rows.Scan(
			&comment.Id,
			&comment.PostId,
			&comment.UserId,
			&comment.Content,
			&comment.Likes,
			&comment.Dislikes,
		)
		if err != nil {
			return nil, err
		}
		response.Comments = append(response.Comments, &comment)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *commentRepoPostgres) GetCommentsByUserId(post_id *pbc.PostId) (*pbc.GetAllCommentsResponse, error) {
	rows, err := c.db.Query(`
		SELECT
			id,
			post_id,
			user_id,
			content,
			likes,
			dislikes
		FROM comments
		WHERE user_id = $1 AND deleted_at IS NULL`, post_id.PostId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var response pbc.GetAllCommentsResponse
	for rows.Next() {
		var comment pbc.Comment
		err := rows.Scan(
			&comment.Id,
			&comment.PostId,
			&comment.UserId,
			&comment.Content,
			&comment.Likes,
			&comment.Dislikes,
		)
		if err != nil {
			return nil, err
		}
		response.Comments = append(response.Comments, &comment)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *commentRepoPostgres) UpdateComment(comment *pbc.Comment) (*pbc.Comment, error) {
	UpdatedAt := time.Now()
	_, err := c.db.Exec(`
		UPDATE comments
			SET
				content = $1,
				likes = $2,
				dislikes = $3,
				updated_at = $4
			WHERE id = $5 AND deleted_at IS NULL`,
		comment.Content,
		comment.Likes,
		comment.Dislikes,
		UpdatedAt,
		comment.Id)
	if err != nil {
		return nil, err
	}
	req := &pbc.CommentId{
		Id: comment.Id,
	}
	resp, err := c.GetComment(req)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *commentRepoPostgres) DeleteComment(id *pbc.CommentId) (*pbc.DeleteResponse, error) {
	resp := &pbc.DeleteResponse{}
	deletedAt := time.Now()
	result, err := c.db.Exec(`
		UPDATE comments
		SET deleted_at = $1
		WHERE id = $2 AND deleted_at IS NULL`, deletedAt, id.Id)
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
