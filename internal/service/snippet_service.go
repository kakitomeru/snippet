package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/kakitomeru/shared/interceptor"
	"github.com/kakitomeru/shared/pagination"
	"github.com/kakitomeru/snippet/internal/model"
	"github.com/kakitomeru/snippet/internal/repository"
)

var ErrSnippetIsPrivate = errors.New("snippet is private")

type SnippetService interface {
	Create(
		ctx context.Context,
		title string,
		content string,
		languageHint string,
		isPublic bool,
		tags []string,
	) (*uuid.UUID, error)
	GetSnippet(
		ctx context.Context,
		id uuid.UUID,
	) (*model.Snippet, error)
	List(
		ctx context.Context,
		pagination pagination.Pagination,
		ownerID string,
		excludeMine bool,
	) ([]*model.Snippet, int64, error)
	ListByOwnerID(
		ctx context.Context,
		pagination pagination.Pagination,
	) ([]*model.Snippet, int64, error)
	Update(ctx context.Context, id uuid.UUID, snippet *model.Snippet) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type snippetService struct {
	repo *repository.Repository
}

func NewSnippetService(repo *repository.Repository) *snippetService {
	return &snippetService{repo: repo}
}

func (s *snippetService) Create(
	ctx context.Context,
	title string,
	content string,
	languageHint string,
	isPublic bool,
	tags []string,
) (*uuid.UUID, error) {
	userID := interceptor.GetUserID(ctx)

	snippet := &model.Snippet{
		OwnerID:      userID,
		Title:        title,
		Content:      content,
		LanguageHint: languageHint,
		IsPublic:     isPublic,
		Tags:         tags,
	}

	id, err := s.repo.Snippet.Create(ctx, snippet)
	if err != nil {
		return nil, err
	}

	return id, nil
}

func (s *snippetService) GetSnippet(
	ctx context.Context,
	id uuid.UUID,
) (*model.Snippet, error) {
	userID := interceptor.GetUserID(ctx)

	snippet, err := s.repo.Snippet.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if !snippet.IsPublic && snippet.OwnerID != userID {
		return nil, ErrSnippetIsPrivate
	}

	return snippet, nil
}

func (s *snippetService) List(
	ctx context.Context,
	pagination pagination.Pagination,
	ownerID string,
	excludeMine bool,
) ([]*model.Snippet, int64, error) {
	userID := interceptor.GetUserID(ctx)

	if ownerID != "" {
		ownerID, err := uuid.Parse(ownerID)
		if err != nil {
			return nil, 0, err
		}

		snippets, total, err := s.repo.Snippet.ListByOwnerID(ctx, ownerID, pagination, true)
		if err != nil {
			return nil, total, err
		}
		return snippets, total, nil
	}

	fmt.Println("excludeMine", excludeMine)
	if excludeMine {
		snippets, total, err := s.repo.Snippet.ListExcludeOwnerID(ctx, pagination, userID)
		if err != nil {
			return nil, total, err
		}
		return snippets, total, nil
	}

	snippets, total, err := s.repo.Snippet.List(ctx, pagination)
	if err != nil {
		return nil, total, err
	}

	return snippets, total, nil
}

func (s *snippetService) ListByOwnerID(
	ctx context.Context,
	pagination pagination.Pagination,
) ([]*model.Snippet, int64, error) {
	userID := interceptor.GetUserID(ctx)

	snippets, total, err := s.repo.Snippet.ListByOwnerID(ctx, userID, pagination, false)
	if err != nil {
		return nil, total, err
	}

	return snippets, total, nil
}

func (s *snippetService) Update(
	ctx context.Context,
	id uuid.UUID,
	snippet *model.Snippet,
) error {
	panic("not implemented")
}

func (s *snippetService) Delete(ctx context.Context, id uuid.UUID) error {
	panic("not implemented")
}
