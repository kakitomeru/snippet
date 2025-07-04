package repository

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/kakitomeru/shared/pagination"
	"github.com/kakitomeru/snippet/internal/model"
	"gorm.io/gorm"
)

var ErrSnippetNotFound = errors.New("snippet not found")

type SnippetRepository interface {
	Create(ctx context.Context, snippet *model.Snippet) (*model.Snippet, error)
	GetByID(ctx context.Context, id uuid.UUID) (*model.Snippet, error)
	ListPublic(
		ctx context.Context,
		pagination pagination.Pagination,
	) ([]*model.Snippet, int64, error)
	ListPublicExcludeOwnerID(
		ctx context.Context,
		pagination pagination.Pagination,
		userID uuid.UUID,
	) ([]*model.Snippet, int64, error)
	ListByOwnerID(
		ctx context.Context,
		ownerID uuid.UUID,
		pagination pagination.Pagination,
		excludePrivate bool,
	) ([]*model.Snippet, int64, error)
	Update(
		ctx context.Context,
		snippet *model.Snippet,
		params map[string]any,
	) (*model.Snippet, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

type snippetRepository struct {
	db *gorm.DB
}

func NewSnippetRepository(db *gorm.DB) *snippetRepository {
	return &snippetRepository{db: db}
}

func (r *snippetRepository) Create(
	ctx context.Context,
	snippet *model.Snippet,
) (*model.Snippet, error) {
	snippet.ID = uuid.New()
	snippet.CreatedAt = time.Now()
	snippet.UpdatedAt = time.Now()

	if err := r.db.WithContext(ctx).Create(snippet).Error; err != nil {
		return nil, err
	}

	var createdSnippet model.Snippet
	if err := r.db.WithContext(ctx).Preload("Owner").First(&createdSnippet, "id = ?", snippet.ID).Error; err != nil {
		return nil, err
	}

	return &createdSnippet, nil
}

func (r *snippetRepository) GetByID(
	ctx context.Context,
	id uuid.UUID,
) (*model.Snippet, error) {
	var snippet model.Snippet
	if err := r.db.WithContext(ctx).Preload("Owner").First(&snippet, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrSnippetNotFound
		}

		return nil, err
	}

	return &snippet, nil
}

func (r *snippetRepository) ListPublic(
	ctx context.Context,
	pagination pagination.Pagination,
) ([]*model.Snippet, int64, error) {
	var total int64

	err := r.db.
		WithContext(ctx).
		Model(&model.Snippet{}).
		Where("is_public = true").
		Count(&total).Error
	if err != nil {
		return nil, total, err
	}

	var snippets []*model.Snippet
	err = r.db.
		WithContext(ctx).
		Preload("Owner").
		Where("is_public = true").
		Order("created_at DESC").
		Offset((pagination.Page - 1) * pagination.Size).
		Limit(pagination.Size).
		Find(&snippets).Error
	if err != nil {
		return nil, total, err
	}

	return snippets, total, nil
}

func (r *snippetRepository) ListPublicExcludeOwnerID(
	ctx context.Context,
	pagination pagination.Pagination,
	userID uuid.UUID,
) ([]*model.Snippet, int64, error) {
	var total int64

	err := r.db.
		WithContext(ctx).
		Model(&model.Snippet{}).
		Where("is_public = true AND owner_id != ?", userID).
		Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	var snippets []*model.Snippet
	err = r.db.
		WithContext(ctx).
		Preload("Owner").
		Where("is_public = true AND owner_id != ?", userID).
		Order("created_at DESC").
		Offset((pagination.Page - 1) * pagination.Size).
		Limit(pagination.Size).
		Find(&snippets).Error
	if err != nil {
		return nil, 0, err
	}

	return snippets, total, nil
}

func (r *snippetRepository) ListByOwnerID(
	ctx context.Context,
	ownerID uuid.UUID,
	pagination pagination.Pagination,
	excludePrivate bool,
) ([]*model.Snippet, int64, error) {
	var total int64

	query := "owner_id = ?"
	if excludePrivate {
		query += " AND is_public = true"
	}

	err := r.db.
		WithContext(ctx).
		Model(&model.Snippet{}).
		Where(query, ownerID).
		Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	var snippets []*model.Snippet
	err = r.db.
		WithContext(ctx).
		Preload("Owner").
		Where(query, ownerID).
		Order("created_at DESC").
		Offset((pagination.Page - 1) * pagination.Size).
		Limit(pagination.Size).
		Find(&snippets).Error
	if err != nil {
		return nil, 0, err
	}

	return snippets, total, nil
}

func (r *snippetRepository) Update(
	ctx context.Context,
	snippet *model.Snippet,
	params map[string]any,
) (*model.Snippet, error) {
	snippet.UpdatedAt = time.Now()

	err := r.db.WithContext(ctx).Model(snippet).Updates(params).Error
	if err != nil {
		return nil, err
	}

	var updatedSnippet model.Snippet
	err = r.db.WithContext(ctx).Preload("Owner").First(&updatedSnippet, "id = ?", snippet.ID).Error
	if err != nil {
		return nil, err
	}

	return &updatedSnippet, nil
}

func (r *snippetRepository) Delete(ctx context.Context, id uuid.UUID) error {
	err := r.db.WithContext(ctx).Delete(&model.Snippet{}, "id = ?", id).Error
	if err != nil {
		return err
	}

	return nil
}
