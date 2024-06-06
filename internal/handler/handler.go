package handler

import (
	"github.com/jakofys/fluid/internal/server"
	"github.com/jakofys/fluid/internal/services"
)

type handler struct {
	linkSrv *services.LinkService
}

func NewHandler(linkSrv *services.LinkService) server.StrictServerInterface {
	return &handler{linkSrv: linkSrv}
}
