package service

import "github.com/kakitomeru/snippet/internal/repository"

type Service struct {
	Snippet SnippetService
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Snippet: NewSnippetService(repo),
	}
}
