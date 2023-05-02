package hottie

import (
	"strings"

	log "github.com/charmbracelet/log"
	"github.com/valyala/fasthttp"
)

type ParsedRequest struct {
	Path        string
	ContentType string
}

func (h *hottie) ParseRequest(ctx *fasthttp.RequestCtx) ParsedRequest {
	var ext string
	path := string(ctx.Path())

	pathParts := strings.Split(path, ".")
	switch l := len(pathParts); {
	case l == 1 || path == "/":
		path = "index.html"
		ext = "html"
	case l > 1:
		ext = pathParts[l-1]
	default:
		log.Fatal("bad path")
	}

	contentType := determineContentType(ext)
	if contentType == "" {
		log.Fatal("unable to determine content type")
	}

	return ParsedRequest{
		Path:        path,
		ContentType: contentType,
	}
}

func determineContentType(extension string) string {
	return ""
}
