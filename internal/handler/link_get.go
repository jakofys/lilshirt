package handler

import (
	"context"

	"github.com/jakofys/fluid/internal/server"
)

func (h *handler) LinkGet(ctx context.Context, request server.LinkGetRequestObject) (server.LinkGetResponseObject, error) {
	link, err := h.linkSrv.GetLink(ctx, request.LinkId)
	if err != nil {
		return nil, err
	}
	return server.LinkGet200JSONResponse{
		LinkGetResponseJSONResponse: server.LinkGetResponseJSONResponse{
			ID:        link.ID,
			Long:      *link.Long,
			Short:     *link.Short,
			CreatedAt: link.CreatedAt,
			ExpiredAt: link.ExpiredAt},
	}, nil
}
