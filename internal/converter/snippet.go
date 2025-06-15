package converter

import (
	"github.com/kakitomeru/snippet/internal/model"
	pb "github.com/kakitomeru/snippet/pkg/pb/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToPbSnippet(snippet *model.Snippet) *pb.Snippet {
	return &pb.Snippet{
		Id:           snippet.ID.String(),
		OwnerId:      snippet.OwnerID.String(),
		Title:        snippet.Title,
		Content:      snippet.Content,
		LanguageHint: snippet.LanguageHint,
		IsPublic:     snippet.IsPublic,
		Tags:         snippet.Tags,
		CreatedAt:    timestamppb.New(snippet.CreatedAt),
		UpdatedAt:    timestamppb.New(snippet.UpdatedAt),
	}
}
