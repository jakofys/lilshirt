package handler

import (
	"context"

	"github.com/jakofys/fluid/internal/server"
)

func (h *handler) LinkCreate(ctx context.Context, request server.LinkCreateRequestObject) (server.LinkCreateResponseObject, error) {
	link, err := h.linkSrv.GenerateLink(ctx, request.Body.Long)
	if err != nil {
		return nil, err
	}
	return server.LinkCreate201JSONResponse{
		server.LinkCreatedResponseJSONResponse{
			ID:        link.ID,
			Long:      *link.Long,
			Short:     *link.Short,
			CreatedAt: link.CreatedAt,
			ExpiredAt: link.ExpiredAt,
		},
	}, nil
}
