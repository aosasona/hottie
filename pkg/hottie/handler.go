package hottie

import (
	"bufio"
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
	ctx.SetContentType("text/event-stream")
	ctx.Response.Header.Set("Cache-Control", "no-cache")
	ctx.Response.Header.Set("Connection", "keep-alive")
	ctx.Response.Header.Set("Transfer-Encoding", "chunked")
	ctx.Response.Header.Set("Access-Control-Allow-Origin", "*")
	ctx.Response.Header.Set("Access-Control-Allow-Headers", "Cache-Control")
	ctx.Response.Header.Set("Access-Control-Allow-Credentials", "true")
	ctx.SetBodyStreamWriter(fasthttp.StreamWriter(func(w *bufio.Writer) {
		for {
			select {
			case e := <-notifChan:
				var reloadType string
				switch e {
				case full_reload:
					reloadType = "_full"
				case css_reload:
					reloadType = "_style"
				}
				w.WriteString("data: " + reloadType + "\n\n")
				w.Flush()
			}
		}
	}))
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
