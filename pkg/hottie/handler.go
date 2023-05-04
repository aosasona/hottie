package hottie

import (
	"os"

	"github.com/valyala/fasthttp"
)

func (h *hottie) serveFile(ctx *fasthttp.RequestCtx) {
	parsedRequest := h.parseRequest(ctx)
	if parsedRequest.FileType == HTML {
		h.handleHTMLRequest(ctx, parsedRequest)
		return
	}

	h.handleOtherRequest(ctx, parsedRequest)
	return
}

func (h *hottie) handleHTMLRequest(ctx *fasthttp.RequestCtx, parsedRequest ParsedRequest) {
	file, errMsg, statusCode := h.getFile(parsedRequest.Path)
	if statusCode != fasthttp.StatusOK {
		ctx.Error(errMsg, statusCode)
		return
	}

	ctx.SetContentType(parsedRequest.ContentType)
	ctx.Response.Header.Set("Cache-Control", "no-cache, no-store, must-revalidate")
	ctx.Response.Header.Set("Access-Control-Allow-Headers", "Cache-Control")
	ctx.SetStatusCode(fasthttp.StatusOK)

	content := injectWebsocketCode(file)

	ctx.SetBody(content)
	return
}

func (h *hottie) handleOtherRequest(ctx *fasthttp.RequestCtx, parsedRequest ParsedRequest) {
	file, errMsg, statusCode := h.getFile(parsedRequest.Path)
	if statusCode != fasthttp.StatusOK {
		ctx.Error(errMsg, statusCode)
		return
	}

	ctx.SetContentType(parsedRequest.ContentType)
	ctx.Response.Header.Set("Cache-Control", "no-cache, no-store, must-revalidate")
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetBody(file)
	return
}

func (h *hottie) handleSSE(ctx *fasthttp.RequestCtx) {
}

func (h *hottie) getFile(path string) ([]byte, string, int) {
	var (
		file []byte
		msg  string
		err  error
	)

	filePath := h.dir + "/" + path
	file, err = os.ReadFile(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return file, "404 page not found", fasthttp.StatusNotFound
		}
		return file, "Internal Server Error", fasthttp.StatusInternalServerError
	}

	return file, msg, fasthttp.StatusOK
}
