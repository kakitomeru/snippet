package api

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/kakitomeru/shared/pagination"
	s "github.com/kakitomeru/shared/status"
	"github.com/kakitomeru/snippet/internal/converter"
	"github.com/kakitomeru/snippet/internal/repository"
	"github.com/kakitomeru/snippet/internal/service"
	pb "github.com/kakitomeru/snippet/pkg/pb/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type SnippetServiceHandler struct {
	pb.UnimplementedSnippetServiceServer
	service *service.Service
}

func NewSnippetServiceHandler(service *service.Service) *SnippetServiceHandler {
	return &SnippetServiceHandler{service: service}
}

func (h *SnippetServiceHandler) CreateSnippet(
	ctx context.Context,
	req *pb.CreateSnippetRequest,
) (*pb.CreateSnippetResponse, error) {
	if req.Content == "" || req.Title == "" {
		return nil, status.Errorf(codes.InvalidArgument, "content and title are required")
	}

	languageHint := "text"
	if req.LanguageHint != nil {
		languageHint = *req.LanguageHint
	}

	isPublic := false
	if req.IsPublic != nil {
		isPublic = *req.IsPublic
	}

	tags := []string{}
	if req.Tags != nil {
		tags = req.Tags
	}

	id, err := h.service.Snippet.Create(ctx, req.Title, req.Content, languageHint, isPublic, tags)
	if err != nil {
		return nil, s.StatusInternal
	}

	return &pb.CreateSnippetResponse{
		Id: id.String(),
	}, nil
}

func (h *SnippetServiceHandler) GetSnippet(
	ctx context.Context,
	req *pb.GetSnippetRequest,
) (*pb.GetSnippetResponse, error) {
	if req.Id == "" {
		return nil, status.Errorf(codes.InvalidArgument, "id is required")
	}

	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid id")
	}

	snippet, err := h.service.Snippet.Get(ctx, id)
	if err != nil {
		switch {
		case errors.Is(err, repository.ErrSnippetNotFound):
			return nil, status.Errorf(codes.NotFound, "snippet not found")
		case errors.Is(err, service.ErrSnippetIsPrivate):
			return nil, status.Errorf(codes.PermissionDenied, "snippet is private")
		default:
			return nil, s.StatusInternal
		}
	}

	return &pb.GetSnippetResponse{
		Snippet: converter.ToPbSnippet(snippet),
	}, nil
}

func (h *SnippetServiceHandler) ListMySnippets(
	ctx context.Context,
	req *pb.ListMySnippetsRequest,
) (*pb.ListMySnippetsResponse, error) {
	page := 1
	if req.Page != nil && *req.Page > 0 {
		page = int(*req.Page)
	}

	size := 10
	if req.Size != nil && *req.Size > 0 && *req.Size <= 50 {
		size = int(*req.Size)
	}

	pagination := pagination.Pagination{
		Page: page,
		Size: size,
	}

	snippets, total, err := h.service.Snippet.ListMine(ctx, pagination)
	if err != nil {
		return nil, s.StatusInternal
	}

	snippetsPb := make([]*pb.Snippet, len(snippets))
	for i, snippet := range snippets {
		snippetsPb[i] = converter.ToPbSnippet(snippet)
	}
	return &pb.ListMySnippetsResponse{
		Snippets:   snippetsPb,
		Pagination: converter.ToPbPagination(pagination, total),
	}, nil
}

func (h *SnippetServiceHandler) ListPublicSnippets(
	ctx context.Context,
	req *pb.ListPublicSnippetsRequest,
) (*pb.ListPublicSnippetsResponse, error) {
	page := 1
	if req.Page != nil && *req.Page > 0 {
		page = int(*req.Page)
	}

	size := 10
	if req.Size != nil && *req.Size > 0 && *req.Size <= 50 {
		size = int(*req.Size)
	}

	pagination := pagination.Pagination{
		Page: page,
		Size: size,
	}

	excludeMine := false
	if req.ExcludeMine != nil {
		excludeMine = *req.ExcludeMine
	}

	var ownerID *uuid.UUID
	if req.OwnerId != nil {
		parsedOwnerID, err := uuid.Parse(*req.OwnerId)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "invalid owner id")
		}
		ownerID = &parsedOwnerID
	}

	snippets, total, err := h.service.Snippet.ListPublic(ctx, pagination, ownerID, excludeMine)
	if err != nil {
		return nil, s.StatusInternal
	}

	snippetsPb := make([]*pb.Snippet, len(snippets))
	for i, snippet := range snippets {
		snippetsPb[i] = converter.ToPbSnippet(snippet)
	}

	return &pb.ListPublicSnippetsResponse{
		Snippets:   snippetsPb,
		Pagination: converter.ToPbPagination(pagination, total),
	}, nil
}

func (h *SnippetServiceHandler) UpdateSnippet(
	ctx context.Context,
	req *pb.UpdateSnippetRequest,
) (*pb.UpdateSnippetResponse, error) {
	if req.Id == "" {
		return nil, status.Errorf(codes.InvalidArgument, "id is required")
	}

	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid id")
	}

	title := ""
	if req.Title != nil {
		title = *req.Title
	}

	content := ""
	if req.Content != nil {
		content = *req.Content
	}

	languageHint := ""
	if req.LanguageHint != nil {
		languageHint = *req.LanguageHint
	}

	var isPublic *bool
	if req.IsPublic != nil {
		isPublic = req.IsPublic
	}

	var tags *[]string
	if req.Tags != nil {
		tags = &req.Tags
	}

	fmt.Println(title, content, languageHint, isPublic, tags, req.Tags)
	fmt.Println("title", req.Title)
	fmt.Println("content", req.Content)
	fmt.Println("languageHint", req.LanguageHint)
	fmt.Println("isPublic", req.IsPublic)
	fmt.Println("tags", req.Tags)

	// h.service.Snippet.Update(ctx, id, title, content, languageHint, isPublic, tags)

	return &pb.UpdateSnippetResponse{
		Snippet: &pb.Snippet{
			Id: id.String(),
		},
	}, nil
}

func (h *SnippetServiceHandler) DeleteSnippet(
	ctx context.Context,
	req *pb.DeleteSnippetRequest,
) (*emptypb.Empty, error) {
	panic("not implemented")
}
