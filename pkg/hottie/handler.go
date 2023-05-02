package hottie

import (
	"fmt"

	"github.com/valyala/fasthttp"
)

func (h *hottie) HandleRequest(ctx *fasthttp.RequestCtx) {
	parsedRequest := h.ParseRequest(ctx)
	path := string(ctx.Path())
	fmt.Println(path)
	fmt.Println(parsedRequest)
}
