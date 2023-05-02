package hottie

import (
	"mime"
	"strings"

	log "github.com/charmbracelet/log"
	"github.com/valyala/fasthttp"
)

const (
	HTML  FileType = "html"
	OTHER FileType = "other"
)

type FileType string

type ParsedRequest struct {
	Path        string
	ContentType string
	FileType    FileType
}

func (h *hottie) ParseRequest(ctx *fasthttp.RequestCtx) ParsedRequest {
	var ext string
	fileType := OTHER

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

	contentType := mime.TypeByExtension("." + ext)
	if contentType == "" {
		log.Fatal("unable to determine content type")
	}

	if strings.Contains(contentType, "text/html") {
		fileType = HTML
	}

	return ParsedRequest{
		Path:        path,
		ContentType: contentType,
		FileType:    fileType,
	}
}
