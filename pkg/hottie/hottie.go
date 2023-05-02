package hottie

import (
	"fmt"
	"os"
	"time"

	"github.com/charmbracelet/log"
	"github.com/valyala/fasthttp"
)

type hottie struct {
	dir             string
	port            int
	enableHotReload bool
	log             *log.Logger
}

func New() *hottie {
	logger := log.NewWithOptions(os.Stdout, log.Options{
		ReportTimestamp: true,
		TimeFormat:      time.Kitchen,
		Prefix:          "Hottie 🔥",
	})
	return &hottie{
		log: logger,
	}
}

func (h *hottie) SetDir(dir string) *hottie {
	h.dir = dir
	return h
}

func (h *hottie) SetPort(port int) *hottie {
	h.port = port
	return h
}

func (h *hottie) SetEnableHotReload(hotReload bool) *hottie {
	h.enableHotReload = hotReload
	return h
}

func (h *hottie) Run() error {
	h.log.Info(
		fmt.Sprintf(
			"Starting Hottie on port %d - serving %s\n\nVisit http://127.0.0.1:%d in your browser\n",
			h.port,
			h.dir,
			h.port,
		),
	)
	return fasthttp.ListenAndServe(fmt.Sprintf(":%d", h.port), h.HandleRequest)
}
