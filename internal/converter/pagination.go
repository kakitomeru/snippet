package converter

import (
	"github.com/kakitomeru/shared/pagination"
	pb "github.com/kakitomeru/snippet/pkg/pb/v1"
)

func ToPbPagination(pag pagination.Pagination, total int64) *pb.PaginationResponse {
	return &pb.PaginationResponse{
		Size:        int32(pag.Size),
		TotalItems:  int32(total),
		CurrentPage: int32(pag.Page),
		TotalPages:  int32(pagination.TotalPages(total, pag.Size)),
	}
}
