package postgre

import (
	"cmd/pkg/storage"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Storage struct {
	db *pgxpool.Pool
}

func NewStorage(db *pgxpool.Pool) *Storage {
	return &Storage{db}
}

// GetAllPosts Возвращает список всех задач
func (s *Storage) GetAllPosts() ([]storage.Post, error) {
	query := `
		SELECT id, author_id, title, content, created_at
		FROM posts
	`
	rows, err := s.db.Query(context.Background(), query)
	if err != nil {
		return nil, fmt.Errorf("GetAllPosts: %s", err)
	}
	defer rows.Close()

	var posts []storage.Post
	for rows.Next() {
		var post storage.Post

		err := rows.Scan(&post.ID, &post.AuthorID, &post.Title, &post.Content, &post.Created)
		if err != nil {
			return nil, fmt.Errorf("GetAllPosts: %s", err)
		}

		posts = append(posts, post)
	}
	return posts, nil
}

// AddPost Создает новую публикацию
func (s *Storage) AddPost(post storage.Post) error {
	query := `
		INSERT INTO posts 
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`

	var id int
	err := s.db.QueryRow(context.Background(), query, post.ID, post.AuthorID, post.Title, post.Content, post.Created).Scan(&id)
	if err != nil {
		return fmt.Errorf("AddPost: %s", err)
	}
	fmt.Printf("Added post with ID: %d\n", id)
	return nil
}

// UpdatePost Обновление публикации
func (s *Storage) UpdatePost(post storage.Post) error {
	query := `
		UPDATE posts
		SET title=$1, content=$2, author_id=$3
	`

	_, err := s.db.Exec(context.Background(), query, post.Title, post.Content, post.AuthorID)
	if err != nil {
		return fmt.Errorf("UpdatePost: %s", err)
	}
	fmt.Printf("Updated post with ID: %d\n", post.ID)
	return nil
}

// DeletePost Удаление публикации по ID
func (s *Storage) DeletePost(post storage.Post) error {
	query := `
		DELETE FROM posts WHERE id = $1
	`

	_, err := s.db.Exec(context.Background(), query, post.ID)
	if err != nil {
		return fmt.Errorf("DeletePost: %s", err)
	}
	fmt.Printf("Deleted post with ID: %d\n", post.ID)
	return nil
}
