package hottie

import (
	"fmt"

	"github.com/valyala/fasthttp"
)

type hottie struct {
	dir  string
	port int
}

func New() *hottie {
	return &hottie{}
}

func (h *hottie) SetDir(dir string) *hottie {
	h.dir = dir
	return h
}

func (h *hottie) SetPort(port int) *hottie {
	h.port = port
	return h
}

func (h *hottie) Run() error {
	return fasthttp.ListenAndServe(fmt.Sprintf(":%d", h.port), h.HandleRequest)
}
