package blog

import (
	"context"
	"homepage/internal/database"
	"homepage/internal/markdown"
	"homepage/internal/models"
)

type BlogService struct {
	DB *database.Queries
}

func NewBlogService(db *database.Queries) *BlogService {
	return &BlogService{DB: db}
}

func (s *BlogService) ListBlogPosts(ctx context.Context) ([]models.Post, error) {
	dbPosts, err := s.DB.ListBlogPosts(ctx)
	if err != nil {
		return nil, err
	}

	// Convert database.Post slice to models.Post slice
	modelPosts := make([]models.Post, len(dbPosts))
	for i, dbPost := range dbPosts {
		modelPosts[i] = convertDatabasePostToModelPost(dbPost)
	}

	return modelPosts, nil
}

func convertDatabasePostToModelPost(dbPost database.Post) models.Post {
	return models.Post{
		ID:      dbPost.ID,
		Title:   dbPost.Title,
		Content: dbPost.Content,
		// Add other fields as necessary
	}
}

func (s *BlogService) GetBlogPost(ctx context.Context, id int32) (models.Post, string, error) {
	dbPost, err := s.DB.GetBlogPost(ctx, id)
	if err != nil {
		return models.Post{}, "", err
	}
	post := models.Post{
		ID:      dbPost.ID,
		Title:   dbPost.Title,
		Content: dbPost.Content,
	}
	htmlContent, err := markdown.ConvertAndSanitize(post.Content)
	if err != nil {
		return models.Post{}, "", err
	}

	return post, htmlContent, nil
}

func (s *BlogService) CreateBlogPost(ctx context.Context, title, content string) (models.Post, error) {
	params := database.CreateBlogPostParams{
		Title:   title,
		Content: content,
	}
	result, err := s.DB.CreateBlogPost(ctx, params)
	if err != nil {
		return models.Post{}, err
	}
	return models.Post{ID: result.ID, Title: title, Content: content}, nil
}
