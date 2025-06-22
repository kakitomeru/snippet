package model

import (
	"github.com/google/uuid"
	auth "github.com/kakitomeru/auth/pkg/model"
	"github.com/kakitomeru/shared/model"
	"github.com/lib/pq"
)

type Snippet struct {
	model.Model

	Title        string         `gorm:"not null"`
	Content      string         `gorm:"not null"`
	LanguageHint string         `gorm:"not null;default:'text'"`
	IsPublic     bool           `gorm:"not null;default:false"`
	Tags         pq.StringArray `gorm:"type:text[]"`
	OwnerID      uuid.UUID      `gorm:"not null"`

	Owner auth.User `gorm:"foreignKey:OwnerID"`
}

func (a *Snippet) Equal(b *Snippet) bool {
	if a == nil || b == nil {
		return a == b
	}

	if a.ID != b.ID ||
		a.OwnerID != b.OwnerID ||
		a.CreatedAt != b.CreatedAt ||
		a.UpdatedAt != b.UpdatedAt ||
		a.Title != b.Title ||
		a.Content != b.Content ||
		a.LanguageHint != b.LanguageHint ||
		a.IsPublic != b.IsPublic {
		return false
	}

	if len(a.Tags) != len(b.Tags) {
		return false
	}

	for i := range a.Tags {
		if a.Tags[i] != b.Tags[i] {
			return false
		}
	}

	return true
}
