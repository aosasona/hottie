package hottie

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/fasthttp/websocket"
	"github.com/valyala/fasthttp"
)

const (
	writeWait  = 10 * time.Second
	pongWait   = 60 * time.Second
	pingPeriod = (pongWait * 9) / 10
	filePeriod = 1 * time.Second
)

var upgrader = websocket.FastHTTPUpgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

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
	ctx.SetStatusCode(fasthttp.StatusOK)

	content := injectWebsocketCode(file, fmt.Sprintf("ws://%s:%d/ws", h.addr, h.port))

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

func (h *hottie) handleWS(ctx *fasthttp.RequestCtx) {
	err := upgrader.Upgrade(ctx, func(conn *websocket.Conn) {
	})
	if err != nil {
		log.Println(err)
		return
	}
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
