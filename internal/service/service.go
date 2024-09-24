package service

import (
	"context"
	"fmt"
	"homepage/internal/database"
	"homepage/internal/models"
	"homepage/internal/utils"
)

type ContentService interface {
	GetContentsByType(ctx context.Context, contentTypeStr string) (models.ContentProps, error)
	GetContentById(ctx context.Context, id int) (models.ContentProps, error)
}

type contentService struct {
	DB *database.Queries
}

func NewContentService(db *database.Queries) ContentService {
	return &contentService{
		DB: db,
	}
}

func ParseContentType(s string) (database.ContentType, error) {
	switch s {
	case "blog", "project", "bio":
		return database.ContentType(s), nil
	default:
		return "", fmt.Errorf("invalid content type: %s", s)
	}
}

func (s *contentService) GetContentsByType(ctx context.Context, contentTypeStr string) (models.ContentProps, error) {
	contentType, err := ParseContentType(contentTypeStr)
	if err != nil {
		return models.ContentProps{}, fmt.Errorf("invalid content type: %w", err)
	}

	contents, err := s.DB.GetContentsByType(ctx, contentType)
	if err != nil {
		return models.ContentProps{}, fmt.Errorf("failed to fetch contents: %w", err)
	}

	isAdmin := utils.IsUserAdmin(ctx) // Assuming you've moved this logic to a utility function

	return models.ContentProps{
		Content:     contents,
		ContentType: contentType,
		IsAdmin:     isAdmin,
	}, nil
}

func (s *contentService) GetContentById(ctx context.Context, id int) (models.ContentProps, error) {
	content, err := s.DB.GetContentById(ctx, int32(id))
	if err != nil {
		return models.ContentProps{}, fmt.Errorf("failed to fetch contents: %w", err)
	}

	isAdmin := utils.IsUserAdmin(ctx) // Assuming you've moved this logic to a utility function

	return models.ContentProps{
		Content:     []database.Content{content},
		ContentType: content.ContentType,
		IsAdmin:     isAdmin,
	}, nil
}
