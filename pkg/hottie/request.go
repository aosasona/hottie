package hottie

import (
	"mime"
	"strings"

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

func (h *hottie) parseRequest(ctx *fasthttp.RequestCtx) ParsedRequest {
	var ext string
	fileType := OTHER

	path := string(ctx.Request.URI().Path())

	pathParts := strings.Split(path, ".")
	if path == "/" {
		path = "index.html"
		ext = "html"
	}
	if l := len(pathParts); l > 1 {
		ext = pathParts[l-1]
	}

	contentType := mime.TypeByExtension("." + ext)

	if strings.Contains(contentType, "text/html") {
		fileType = HTML
	}

	return ParsedRequest{
		Path:        path,
		ContentType: contentType,
		FileType:    fileType,
	}
}
