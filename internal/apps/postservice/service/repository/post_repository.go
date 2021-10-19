package repository

import (
	"fmt"
	"time"

	"github.com/HotPotatoC/go-microservice/pkg/database"
)

var (
	ErrPostDoesNotExist = fmt.Errorf("post does not exist")
)

// Post defines the post model
type Post struct {
	ID        string `json:"id"`
	AuthorID  string `json:"author_id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	db database.SQL
}

type PostRepository interface {
	Create(authorID, title, content string) (*Post, error)
	Find(id string) (*Post, error)
	Update(id string, title, content string) (*Post, error)
	Delete(id string) error
	List(before time.Time, limit int32) ([]*Post, error)
}

// NewPostRepository returns a new post query repository
func NewPostRepository(db database.SQL) PostRepository {
	return &Post{db: db}
}

const createPostQuery = `
INSERT INTO posts (
	id,
	author_id,
	title,
	content,
	created_at,
	updated_at
)
VALUES (
	$1,
	$2,
	$3,
	$4,
	$5,
	$6
)`

// Create creates a new post
func (p *Post) Create(authorID, title, content string) (*Post, error) {
	postID, err := GenerateID()
	if err != nil {
		return nil, fmt.Errorf("generate post id error: %v", err)
	}

	post := &Post{
		ID:        postID,
		AuthorID:  authorID,
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_, err = p.db.Exec(createPostQuery,
		post.ID, post.AuthorID, post.Title, post.Content, post.CreatedAt, post.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("create post error: %v", err)
	}

	return post, nil
}

const findPostQuery = `
SELECT
	id,
	author_id,
	title,
	content,
	created_at,
	updated_at
FROM posts
WHERE id = $1
LIMIT 1
`

// Find finds a post by id
func (p *Post) Find(id string) (*Post, error) {
	var post Post
	err := p.db.QueryRow(findPostQuery, id).
		Scan(&post.ID, &post.AuthorID, &post.Title, &post.Content, &post.CreatedAt, &post.UpdatedAt)
	if err != nil {
		return nil, ErrPostDoesNotExist
	}

	return &post, nil
}

const updatePostQuery = `
UPDATE posts SET
	title = $1,
	content = $2,
	updated_at = $3
WHERE id = $4
`

// Update modifies an existing post
func (p *Post) Update(id string, title, content string) (*Post, error) {
	post, err := p.Find(id)
	if err != nil {
		return nil, err
	}

	post.Title = title
	post.Content = content
	post.UpdatedAt = time.Now()

	_, err = p.db.Exec(updatePostQuery,
		post.Title, post.Content, post.UpdatedAt, id)
	if err != nil {
		return nil, fmt.Errorf("update post error: %v", err)
	}

	return post, nil
}

const deletePostQuery = `
DELETE FROM posts
WHERE id = $1
`

// Delete removes a post record
func (p *Post) Delete(id string) error {
	_, err := p.Find(id)
	if err != nil {
		return err
	}

	_, err = p.db.Exec(deletePostQuery, id)
	if err != nil {
		return fmt.Errorf("delete post error: %v", err)
	}

	return nil
}

const listPostsQuery = `
SELECT
	id,
	author_id,
	title,
	content,
	created_at,
	updated_at
FROM posts
WHERE created_at < $1
ORDER BY created_at DESC
LIMIT $2
`

// List returns a list of posts before the given time with a limit
func (p *Post) List(before time.Time, limit int32) ([]*Post, error) {
	rows, err := p.db.Query(listPostsQuery, before, limit)
	if err != nil {
		return nil, fmt.Errorf("list posts error: %v", err)
	}

	posts := []*Post{}
	for rows.Next() {
		post := &Post{}
		err := rows.Scan(&post.ID, &post.AuthorID, &post.Title, &post.Content, &post.CreatedAt, &post.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("scan post error: %v", err)
		}

		posts = append(posts, post)
	}

	return posts, nil
}
