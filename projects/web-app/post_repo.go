package main

import (
	"database/sql"
	"fmt"
	"net/url"
	"strings"

	"github.com/dromara/carbon/v2"
)

type PostRepository interface {
	CreatePost(title, url string, userId int) (int, error)
	AddComment(userId, postId int, body string) (int, error)
	AddVote(userId, postId int) error
	GetAll(filter Filter) ([]Post, Metadata, error)
	GetById(id int) (*Post, error)
	GetComment(postId int) ([]Comment, error)
}

type SQLPostRepository struct {
	db *sql.DB
}

func NewSQLPostRepository(db *sql.DB) *SQLPostRepository {
	return &SQLPostRepository{
		db: db,
	}
}

func (r *SQLPostRepository) CreatePost(title, url string, userId int) (int, error) {
	stmt, err := r.db.Prepare(`INSERT INTO posts (title, url, user_id) VALUES (?,?,?)`)
	if err != nil {
		return 0, err
	}
	res, err := stmt.Exec(title, url, userId)
	if err != nil {
		if strings.Contains(err.Error(), "UNIQUE constraint failed: post.title") {
			return 0, ErrDuplicatePostTitle
		}
		return 0, err
	}

	postId, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(postId), nil
}

func (r *SQLPostRepository) AddComment(userId, postId int, body string) (int, error) {
	stmt, err := r.db.Prepare(`INSERT INTO comments (user_id, post_id, body ) VALUES (?,?,?)`)
	if err != nil {
		return 0, err
	}
	res, err := stmt.Exec(userId, postId, body)
	if err != nil {
		return 0, err
	}

	commentId, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(commentId), nil
}

func (r *SQLPostRepository) AddVote(userId, postId int) error {

	stmt, err := r.db.Prepare(`INSERT INTO votes (user_id, post_id ) VALUES (?,?)`)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(userId, postId)
	if err != nil {
		if strings.Contains(err.Error(), "UNIQUE Constraint failed") || strings.Contains(err.Error(), "PRIMARY KEY constraint failed") {
			return ErrDuplicateVote
		}
		return err
	}
	return nil
}

func (r *SQLPostRepository) GetAll(filter Filter) ([]Post, Metadata, error) {
	if err := filter.Validate(); err != nil {
		return nil, Metadata{}, err
	}

	baseQuery := `
		SELECT 
			COUNT(*) OVER() as total_records,
			p.id, p.title, p.url, p.user_id, p.created_at,
			u.name as user_name,
			COUNT(DISTINCT c.id) as comment_count,
			COUNT(DISTINCT v.user_id) as vote_count
		FROM posts p
		LEFT JOIN users u ON p.user_id = u.id
		LEFT JOIN comments c ON p.id = c.post_id
		LEFT JOIN votes v ON p.id = v.post_id
	`
	var args []interface{}
	if filter.Query != "" {
		baseQuery += " WHERE LOWER(p.title) LIKE ?"
		args = append(args, "%"+strings.ToLower(filter.Query)+"%")
	}
	baseQuery += " GROUP BY p.id, p.title, p.url, p.user_id, p.created_at, u.name"
	if filter.OrderBy == "popular" {
		baseQuery += " ORDER BY vote_count DESC, p.created_at DESC"
	} else {
		baseQuery += " ORDER BY p.created_at DESC"
	}

	limit := filter.PageSize
	offset := (filter.Page - 1) * filter.PageSize
	baseQuery += " LIMIT ? OFFSET ?"
	args = append(args, limit, offset)

	rows, err := r.db.Query(baseQuery, args...)
	if err != nil {
		return nil, Metadata{}, err
	}

	defer rows.Close()

	var posts []Post
	var totalRecords int
	for rows.Next() {
		var post Post
		err := rows.Scan(&totalRecords, &post.ID, &post.Title, &post.URL, &post.UserID,
			&post.CreatedAt, &post.UserName, &post.CommentCount, &post.VoteCount)
		if err != nil {
			return nil, Metadata{}, err
		}
		post.TotalRecords = totalRecords
		posts = append(posts, post)

	}

	if err := rows.Err(); err != nil {
		return nil, Metadata{}, err
	}

	if len(posts) == 0 {
		return []Post{}, Metadata{}, nil
	}

	metadata := calculateMetaData(totalRecords, filter.Page, filter.PageSize)
	return posts, metadata, nil
}

func (r *SQLPostRepository) GetById(id int) (*Post, error) {
	query := `
		SELECT p.id, p.title, p.url, p.user_id, p.created_at,
		u.name as user_name,
		COUNT(DISTINCT c.id) AS comment_count
		COUNT(DISTINCT v.user_id) AS vote_count
		FROM posts p
		LEFT JOIN users u ON p.user_id = u.id
		LEFT JOIN comments c ON p.id = c.post_id
		LEFT JOIN votes v ON p.id = c.post_id
		WHERE p.id = ?
		GROUP BY p.id, p.title, p.url, p.user_id, p.created_at, u.name
	`

	row := r.db.QueryRow(query, id)
	var post Post
	err := row.Scan(
		&post.ID,
		&post.Title,
		&post.URL,
		&post.UserID,
		&post.CreatedAt,
		&post.UserName,
		&post.CommentCount,
		&post.VoteCount,
	)

	if err != nil {
		return nil, err
	}

	return &post, nil
}

func (r *SQLPostRepository) GetComment(postId int) ([]Comment, error) {
	query := `
		SELECT c.id, c.body, c.user_id, c.post_id, c.created_at, u.name as user_name
		FROM comments c
		LEFT JOIN users u ON c.user_id = u.id
		WHERE c.post_id = ?
		ORDER BY c.created_at ASC
	`
	rows, err := r.db.Query(query, postId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var comments []Comment
	for rows.Next() {
		var comment Comment
		err := rows.Scan(
			&comment.ID,
			&comment.Body,
			&comment.UserID,
			&comment.PostID,
			&comment.CreatedAt,
			&comment.UserName,
		)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	if len(comments) == 0 {
		return []Comment{}, nil
	}

	return comments, nil

}

func (p *Post) GetVoteCountHuman() string {
	if p.VoteCount > 1 {
		return fmt.Sprintf("%d votes count", p.VoteCount)
	}
	return fmt.Sprintf("%d vote count", p.VoteCount)
}

func (p *Post) GetCommentCountHuman() string {
	if p.CommentCount > 1 {
		return fmt.Sprintf("%d comments count", p.CommentCount)
	}
	return fmt.Sprintf("%d comment count", p.CommentCount)
}

func (p *Post) CreatedAtHuman() string {
	return carbon.NewCarbon(p.CreatedAt).DiffForHumans()
}

func (p *Post) Host() string {
	ur, err := url.Parse(p.URL)
	if err != nil {
		return "<invalid-host>"
	}
	return ur.Hostname()
}
