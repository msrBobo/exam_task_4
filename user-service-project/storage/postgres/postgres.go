package postgres

import (
	pbu "exam_task_4/user-service-project/genproto/user-service"
	"time"

	"github.com/jmoiron/sqlx"
)

type userRepoPostgres struct {
	db *sqlx.DB
}

func NewUserRepoPostgres(db *sqlx.DB) *userRepoPostgres {
	return &userRepoPostgres{db: db}
}

func (u *userRepoPostgres) CreateUser(user *pbu.User) (*pbu.User, error) {
	respUser := &pbu.User{}
	CreatedAt := time.Now()
	err := u.db.QueryRow(`
		INSERT INTO 
			users(
				id,
				username,
				email,
				password,
				created_at,
				first_name,
				last_name,
				bio,
				website
			)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
			RETURNING
				id,
				username,
				email,
				password,
				created_at,
				first_name,
				last_name,
				bio,
				website`,
		user.Id,
		user.Username,
		user.Email,
		user.Password,
		CreatedAt,
		user.FirstName,
		user.LastName,
		user.Bio,
		user.Website,
	).Scan(
		&respUser.Id,
		&respUser.Username,
		&respUser.Email,
		&respUser.Password,
		&respUser.CreatedAt,
		&respUser.FirstName,
		&respUser.LastName,
		&respUser.Bio,
		&respUser.Website,
	)
	if err != nil {
		return nil, err
	}

	return respUser, nil
}

func (u *userRepoPostgres) GetUser(id *pbu.UserId) (*pbu.User, error) {
	var user pbu.User
	err := u.db.QueryRow(`
	SELECT
			id,
			username,
			email,
			password,
			first_name,
			last_name,
			bio,
			website
            FROM users
			WHERE id = $1 AND deleted_at IS NULL`,
		id.UserId,
	).Scan(
		&user.Id,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.FirstName,
		&user.LastName,
		&user.Bio,
		&user.Website,
	)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *userRepoPostgres) GetAllUser(req *pbu.GetAllUsersRequest) (*pbu.GetAllUsersResponse, error) {
	offset := (req.Page - 1) * req.Limit
	rows, err := u.db.Query(`
				SELECT 
					id,
					username,
					email,
					password,
					first_name,
					last_name,
					bio,
					website
						FROM users
					WHERE deleted_at IS NULL
					LIMIT $1 OFFSET $2`, req.Limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var response pbu.GetAllUsersResponse
	for rows.Next() {
		var user pbu.User
		err := rows.Scan(
			&user.Id,
			&user.Username,
			&user.Email,
			&user.Password,
			&user.FirstName,
			&user.LastName,
			&user.Bio,
			&user.Website,
		)
		if err != nil {
			return nil, err
		}
		response.Users = append(response.Users, &user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &response, nil
}

func (u *userRepoPostgres) UpdateUser(user *pbu.User) (*pbu.User, error) {

	UpdatedAt := time.Now()
	_, err := u.db.Exec(`
		UPDATE users
			SET
				username = $1,
				email = $2,
				password = $3,
				updated_at = $4,
				first_name = $5,
				last_name = $6,
				bio = $7,
				website = $8
					WHERE id = $9 AND deleted_at IS NULL`,
		user.Username,
		user.Email,
		user.Password,
		UpdatedAt,
		user.FirstName,
		user.LastName,
		user.Bio,
		user.Website,
		user.Id)
	if err != nil {
		return nil, err
	}
	userId := &pbu.UserId{
		UserId: user.Id,
	}

	resp, err := u.GetUser(userId)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (u *userRepoPostgres) DeleteUser(id *pbu.UserId) (resp *pbu.DeleteResponse, err error) {
	resp = &pbu.DeleteResponse{}
	deletedAt := time.Now()
	result, err := u.db.Exec(`UPDATE users
                            SET deleted_at = $1
                            WHERE id = $2 AND deleted_at IS NULL`, deletedAt, id.UserId)
	if err != nil {
		return nil, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rowsAffected == 0 {
		resp.Message = "User not found or already deleted"
		return resp, nil
	}
	resp.Message = "Deleted User"

	return resp, nil
}
