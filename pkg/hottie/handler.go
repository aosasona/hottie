package hottie

import (
	"os"

	"github.com/valyala/fasthttp"
)

func (h *hottie) HandleRequest(ctx *fasthttp.RequestCtx) {
	parsedRequest := h.ParseRequest(ctx)
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
	ctx.SetStatusCode(fasthttp.StatusOK)

	content := injectWebsocketCode(file, h.port)

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
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetBody(file)
	return
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
