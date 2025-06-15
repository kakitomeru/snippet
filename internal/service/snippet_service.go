package service

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/kakitomeru/shared/interceptor"
	"github.com/kakitomeru/shared/pagination"
	"github.com/kakitomeru/snippet/internal/model"
	"github.com/kakitomeru/snippet/internal/repository"
)

var ErrSnippetIsPrivate = errors.New("snippet is private")
var ErrSnippetNotOwned = errors.New("unauthorized snippet action")

type SnippetService interface {
	Create(
		ctx context.Context,
		title string,
		content string,
		languageHint string,
		isPublic bool,
		tags []string,
	) (*model.Snippet, error)
	Get(
		ctx context.Context,
		id uuid.UUID,
	) (*model.Snippet, error)
	ListPublic(
		ctx context.Context,
		pagination pagination.Pagination,
		ownerID *uuid.UUID,
		excludeMine bool,
	) ([]*model.Snippet, int64, error)
	ListMine(
		ctx context.Context,
		pagination pagination.Pagination,
	) ([]*model.Snippet, int64, error)
	Update(
		ctx context.Context,
		id uuid.UUID,
		title *string,
		content *string,
		languageHint *string,
		isPublic *bool,
		tags []string,
		tagsUpdated bool,
	) (*model.Snippet, error)
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
) (*model.Snippet, error) {
	userID := interceptor.GetUserID(ctx)

	snippet := &model.Snippet{
		OwnerID:      userID,
		Title:        title,
		Content:      content,
		LanguageHint: languageHint,
		IsPublic:     isPublic,
		Tags:         tags,
	}

	snippet, err := s.repo.Snippet.Create(ctx, snippet)
	if err != nil {
		return nil, err
	}

	return snippet, nil
}

func (s *snippetService) Get(
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

func (s *snippetService) ListPublic(
	ctx context.Context,
	pagination pagination.Pagination,
	ownerID *uuid.UUID,
	excludeMine bool,
) ([]*model.Snippet, int64, error) {
	userID := interceptor.GetUserID(ctx)

	if ownerID != nil {
		if *ownerID == userID && excludeMine {
			return []*model.Snippet{}, 0, nil
		}

		snippets, total, err := s.repo.Snippet.ListByOwnerID(ctx, *ownerID, pagination, true)
		if err != nil {
			return nil, total, err
		}
		return snippets, total, nil
	}

	if excludeMine {
		snippets, total, err := s.repo.Snippet.ListPublicExcludeOwnerID(ctx, pagination, userID)
		if err != nil {
			return nil, total, err
		}
		return snippets, total, nil
	}

	snippets, total, err := s.repo.Snippet.ListPublic(ctx, pagination)
	if err != nil {
		return nil, total, err
	}

	return snippets, total, nil
}

func (s *snippetService) ListMine(
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
	title *string,
	content *string,
	languageHint *string,
	isPublic *bool,
	tags []string,
	tagsUpdated bool,
) (*model.Snippet, error) {
	userID := interceptor.GetUserID(ctx)

	snippet, err := s.repo.Snippet.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if snippet.OwnerID != userID {
		return nil, ErrSnippetNotOwned
	}

	updatedSnippet := &model.Snippet{
		OwnerID: snippet.OwnerID,
	}
	updatedSnippet.ID = snippet.ID
	updatedSnippet.CreatedAt = snippet.CreatedAt
	updatedSnippet.UpdatedAt = snippet.UpdatedAt

	if title != nil {
		updatedSnippet.Title = *title
	} else {
		updatedSnippet.Title = snippet.Title
	}

	if content != nil {
		updatedSnippet.Content = *content
	} else {
		updatedSnippet.Content = snippet.Content
	}

	if languageHint != nil {
		updatedSnippet.LanguageHint = *languageHint
	} else {
		updatedSnippet.LanguageHint = snippet.LanguageHint
	}

	if isPublic != nil {
		updatedSnippet.IsPublic = *isPublic
	} else {
		updatedSnippet.IsPublic = snippet.IsPublic
	}

	if tagsUpdated {
		updatedSnippet.Tags = tags
	} else {
		updatedSnippet.Tags = snippet.Tags
	}

	if snippet.Equal(updatedSnippet) {
		return snippet, nil
	}

	updatedSnippet, err = s.repo.Snippet.Update(ctx, id, updatedSnippet)
	if err != nil {
		return nil, err
	}

	return updatedSnippet, nil
}

func (s *snippetService) Delete(ctx context.Context, id uuid.UUID) error {
	userID := interceptor.GetUserID(ctx)

	snippet, err := s.repo.Snippet.GetByID(ctx, id)
	if err != nil {
		return err
	}

	if snippet.OwnerID != userID {
		return ErrSnippetNotOwned
	}

	err = s.repo.Snippet.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
